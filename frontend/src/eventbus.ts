import { Subject } from "rxjs"
import { createLogger } from "./logger"
import { ChatroomMessage } from "./types"

const logger = createLogger("eventbus")

class Eventbus<T> {
	private subjects: Map<any, any> = new Map()

	public publish<K extends keyof T>(event: K, payload: T[K]): void {
		let subj = this.subjects.get(event)

		if (!subj) {
			subj = new Subject<T[K]>()
			this.subjects.set(event, subj)
		}

		if (payload instanceof Array) {
			logger.debug(`publish event: ${String(event)} count: ${payload.length}`)
		} else {
			logger.debug(`publish event: ${String(event)}`)
		}


		subj.next(payload)
	}

	public events<K extends keyof T>(event: K): Subject<T[K]> {
		let subj = this.subjects.get(event)

		if (!subj) {
			subj = new Subject<T[K]>()
			this.subjects.set(event, subj)
		}

		return subj
	}

	public waitForEvent<K extends keyof T>(args: {
		eventType: K
		predicate: (payload: T[K]) => boolean
		timeout?: number
	}): Promise<T[K]> {
		return new Promise((resolve, reject) => {
			const tim = setTimeout(() => {
				sub.unsubscribe()
				reject(new Error("timeout"))
			}, args.timeout || 3_000)

			const sub = this.events(args.eventType).subscribe(payload => {
				if (!args.predicate(payload)) {
					return
				}

				sub.unsubscribe()
				clearTimeout(tim)
				resolve(payload)
			})
		})
	}
}

export type EventMap = {
	"chatroomMessages": ChatroomMessage[]
}

export const eventbus = new Eventbus<EventMap>()