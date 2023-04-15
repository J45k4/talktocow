import { cacheBuilder } from "../src/cache"

it("Upsert user and get user", () => {
	const cache = cacheBuilder({
		user: {

		}
	})

	cache.user.upsert({
		id: "1",
		name: "John"
	})

	const item = cache.user.get("1").get()

	expect(item).toEqual({
		id: "1",
		name: "John"
	})
})

it("Get users", () => {
	const cache = cacheBuilder({
		users: {
			array: true
		}
	})

	cache.users.add({
		id: "1",
		name: "John"
	})

	const item = cache.users.get()

	expect(item).toEqual([{
		id: "1",
		name: "John"
	}])
})

it("Update entity with reference to another and get all values", () => {
	const cache = cacheBuilder({
		myChatroom: {
			array: true,
			reference: "chatroom"
		},
		chatroom: {

		}
	})

	cache.myChatroom.add({
		id: "1",
		name: "Chatroom 1",
	})

	const myChatrooms = cache.myChatroom.get()

	expect(myChatrooms).toEqual([{
		id: "1",
		name: "Chatroom 1",
	}])
})

it("Subscribe to entity changes", async () => {
	const cache = cacheBuilder({
		user: {}
	})

	const user = cache.user.get("1")

	let event
	
	const sub = user.sub((v) => {
		event = v
		sub.unsubscribe()
	})

	cache.user.upsert({
		id: "1",
		name: "John"
	})

	expect(event).toEqual({
		id: "1",
		name: "John"
	})
})

it("Entity has member which is array", () => {
	const cache = cacheBuilder({
		chatroom: {
			fields: {
				members: {
					reference: "user",
					array: true
				}
			}
		},
		user: {}
	})

	cache.chatroom.upsert({
		id: "1",
		name: "Chatroom 1",
	})

	const chatroom = cache.chatroom.get("1")

	console.log("chatroom", chatroom)

	chatroom.members.add({
		id: "1",
		name: "John"
	})

	const memebers = chatroom.members.get()

	expect(memebers).toEqual([{
		id: "1",
		name: "John"
	}])
})

it("Subscribe to array changes", async () => {
	const cache = cacheBuilder({
		users: {
			array: true
		}
	})

	const p = new Promise((resolve) => {
		const sub = cache.users.sub((v) => {
			resolve(v)
			sub.unsubscribe()
		})
	})

	cache.users.add({
		id: "1",
		name: "John"
	})

	const val = await p

	expect(val).toEqual([{
		id: "1",
		name: "John"
	}])
})

it("Subscribe to array changes and entity references another entity", async () => {
	const cache = cacheBuilder({
		myChatrooms: {
			array: true,
			reference: "chatroom"
		},
		chatroom: {}
	})

	const p = new Promise((resolve) => {
		const sub = cache.myChatrooms.sub((v) => {
			resolve(v)
			sub.unsubscribe()
		})
	})

	cache.myChatrooms.add({
		id: "1",
		name: "Chatroom 1",
	})

	const val = await p

	expect(val).toEqual([{
		id: "1",
		name: "Chatroom 1",
	}])
})

it(`Subscribe to array changes and entity 
references another entity and array is not empty`, async () => {
	const cache = cacheBuilder({
		myChatrooms: {
			array: true,
			reference: "chatroom"
		},
		chatroom: {}
	})

	cache.myChatrooms.add({
		id: "1",
		name: "Chatroom 1",
	})

	const p = new Promise((resolve) => {
		const sub = cache.myChatrooms.sub((v) => {
			resolve(v)
			sub.unsubscribe()
		})
	})

	cache.chatroom.upsert({
		id: "1",
		name: "qwert",
	})

	const val = await p

	expect(val).toEqual([{
		id: "1",
		name: "qwert",
	}])
})

it("Unsubscribe from entity changes", async () => {
	const cache = cacheBuilder({
		user: {}
	})

	const user = cache.user.get("1")

	let event
	
	const sub = user.sub((v) => {
		event = v
	})
	sub.unsubscribe()

	cache.user.upsert({
		id: "1",
		name: "John"
	})

	expect(event).toBeUndefined()
})

it("Unsubscribe from array changes", async () => {
	const cache = cacheBuilder({
		users: {
			reference: "user",
			array: true
		},
		user: {}
	})

	let event

	const unsub = cache.users.sub((v) => {
		event = v
	})
	unsub()

	cache.users.add({
		id: "1",
		name: "John"
	})

	expect(event).toBeUndefined()
})