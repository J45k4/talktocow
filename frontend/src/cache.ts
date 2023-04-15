import { useEffect, useState } from "react"
import { Chatroom, User } from "./types"
import { getJson } from "./api-methods"
import { createLogger } from "./logger"

const logger = createLogger("cache")

type ReferenceEntity = {
	reference: string
}

type ArrayInitialFetch = {
	array: true
	initialFetch?: () => Promise<any>
}

type SingleInitialFetch = {
	array?: false
	initialFetch?: (id: string) => Promise<any>
}

type FieldsEntity = {
	fields?: {
		[key: string]: EntityInstruction
	}
}

type EntityInstruction = (ArrayInitialFetch | SingleInitialFetch)
	& (FieldsEntity | ReferenceEntity)
	& {
		default?: any
	}

type Instructions = {
	[key: string]: EntityInstruction
}

type NodeWithId = {
	id: string
}

type Sub<T> = {
	sub: (cb: (value: T) => void) => () => void
}

type ArrayNode<T extends NodeWithId> = {
	add: (p: T) => void
	addMany: (p: T[]) => void
} & CacheNode<T[]>

type SingleNode<T extends NodeWithId> = {
	get: (id: string) => CacheNode<T>
	upsert: (p: T) => void
}

type CacheNode<T> = {
	get: () => T
} & Sub<T>

type Cache = {
	myChatrooms: ArrayNode<Chatroom>
	chatroom: SingleNode<Chatroom>
}

class Subscriptions {
	private map = new Map<string, Set<any>>()

	public sub(args: {
		entityId: string
		next: (value: any) => void
		cleanup?: () => void
	}) {
		let s = this.map.get(args.entityId)

		if (!s) {
			s = new Set()
			this.map.set(args.entityId, s)
		}

		s.add(args.next)

		return () => {
			if (args.cleanup) {
				args.cleanup()
			}

			s.delete(args.next)

			if (s.size === 0) {
				this.map.delete(args.entityId)
			}
		}
	}

	public pub(entityId: string, value: any) {
		const s = this.map.get(entityId)

		if (s) {
			for (const sub of s) {
				sub(value)
			}
		}
	}
}

const createEntityId = (entityType: string, id: string) => `${entityType}:${id}`

export const cacheBuilder = (instructions: Instructions): any => {
	const entityMap = new Map<string, any>()
	const subscriptions = new Subscriptions()

	const upsertEntity = (entityType: string, id: string, data: any) => {
		const entityId = createEntityId(entityType, id)
		entityMap.set(entityId, data)
		subscriptions.pub(entityId, data)
	}

	const processInstruction = (key: string, instruction: EntityInstruction) => {
		const obj: any = {}

		if (instruction.array) {
			obj.sub = (cb: (v) => void) => {
				logger.debug("array sub", key)

				if ("reference" in instruction) {
					const subs = new Map<string, () => void>()

					const puball = () => {
						const refs = entityMap.get(key)
						const m = refs.map((id) => entityMap.get(id))
						cb(m)
					}

					const addsub = (ref) => {
						const unsub = subscriptions.sub({
							entityId: ref,
							next: (data) => {
								puball()
							}
						})
						subs.set(ref, unsub)
					}

					if (entityMap.has(key)) {
						for (const ref of entityMap.get(key)) {
							addsub(ref)
						}
					}

					return subscriptions.sub({
						entityId: key,
						next: (references) => {
							logger.debug("sub changed refs", references)

							const refs = new Set(references)
							
							for (const ref of references) {
								if (subs.has(ref)) {
									continue
								}

								addsub(ref)
							}

							for (const [ref, unsub] of subs) {
								if (refs.has(ref)) {
									continue
								}

								unsub()
								subs.delete(ref)
							}

							puball()
						},
						cleanup: () => {
							for (const [,unsub] of subs) {
								unsub()
							}
						}
					})
				}

				return subscriptions.sub({
					entityId: key,
					next: cb
				})
			}

			obj.get = () => {
				logger.debug("array get", key)

				if (!entityMap.has(key)) {
					if (instruction.initialFetch) {
						logger.debug("instruction has initial fetch", key)

						instruction.initialFetch().then((data) => {
							logger.debug("initial fetch done", key, data)

							if ("reference" in instruction) {
								const entityIds = data.map(p => createEntityId(instruction.reference, p.id))
								entityMap.set(key, entityIds)

								logger.debug("setting entity ids", key, entityIds)

								for (const p of data) {
									const entityId = createEntityId(instruction.reference, p.id)
									entityMap.set(entityId, p)

									logger.debug("setting entity", entityId, p)
								}

								subscriptions.pub(key, entityIds)
								return
							}

							entityMap.set(key, data)
							subscriptions.pub(key, data)
						})
					}

					if (instruction.default) {
						return instruction.default
					}
				}

				const val = entityMap.get(key)

				if ("reference" in instruction) {
					return val.map((id) => entityMap.get(id))
				}

				return val
			}

			obj.add = (p: any) => {
				logger.debug("array add", key, p)

				if ("reference" in instruction) {
					const entityId = createEntityId(instruction.reference, p.id)

					let val = entityMap.get(key)

					if (!val) {
						val = []
						entityMap.set(key, val)	
					}

					val.push(entityId)
					entityMap.set(entityId, p)

					subscriptions.pub(key, val)
					subscriptions.pub(entityId, p)

					return
				}

				let val = entityMap.get(key)

				if (!val) {
					val = []
					entityMap.set(key, val)
				}

				val.push(p)
				subscriptions.pub(key, val)
			}
		} else {
			obj.get = (id: string) => {
				logger.debug("get", key, id)

				const entityId = createEntityId(key, id)
				const entity = entityMap.get(entityId)

				if (!entity && instruction.initialFetch) {
					instruction.initialFetch(id).then((data) => {
						entityMap.set(entityId, entity)
						subscriptions.pub(entityId, entity)
					})
				}

				const returnObj = {
					get: () => entity,
					sub: (cb: (v) => void) => {
						return subscriptions.sub({
							entityId: entityId,
							next: (data) => {
								cb(data)
							}
						})
					}
				}

				if ("fields" in instruction) {
					for (const [key, field] of Object.entries(instruction.fields)) {
						returnObj[key] = processInstruction(key, field)
					}
				}

				return returnObj
			}

			obj.upsert = (p: any) => {
				logger.debug("upsert", key, p)

				if ("reference" in instruction) {
					const ent = entityMap.set(key, p.id)

					const entityId = createEntityId(instruction.reference, p.id)

					entityMap.set(entityId, p)
					subscriptions.pub(entityId, p)

					return
				}

				const entityId = createEntityId(key, p.id)

				entityMap.set(entityId, p)
				subscriptions.pub(entityId, p)
			}
		}

		return obj
	}

	const obj = {}

	for (const [key, instruction] of Object.entries(instructions)) {
		obj[key] = 	processInstruction(key, instruction)
	}

	return obj
}

export const cache: Cache = cacheBuilder({
	"myChatrooms": {
		reference: "chatroom",
		array: true,
		default: [],
		initialFetch: () => {
			return getJson<Chatroom[]>("/api/mychatrooms").then(res => res.payload)
		}
	},
	"user": {
		initialFetch: (id: string) => {
			return getJson<User>(`/api/user/${id}`).then(res => res.payload)
		}
	},
	"users": {
		reference: "user",
		array: true,
		initialFetch: () => {
			return getJson<User[]>("/api/users").then(res => res.payload)
		}
	},
	"chatroom": {
		fields: {
			members: {
				reference: "user",
				array: true,
				// initialFetch: () => {
				// 	return getJson<User[]>(`/api/chatroom/1/members`).then(res => res.payload)
				// }
			}
		},
		// initialFetch: (id: string) => {
		// 	return getJson<Chatroom>(`/api/chatroom/${id}`).then(res => res.payload)
		// }
	}
})

export const useCache = <T>(key: CacheNode<T>) => {
	const [value, setValue] = useState(key.get())

	useEffect(() => {
		const unsub = key.sub(v => {
			logger.debug("sub", key, v)

			setValue(v)
		})

		return () => {
			unsub()
		}
	}, [key])

	return value
}