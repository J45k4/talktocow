

// export class GraphNode {
// 	public constructor(args: {
// 		fields: {
// 			[key: string]: GraphNode
// 		}
// 	}) {

// 	}
// }

function GraphNode<T, M, V>(args: {
	fields: M,
	create?: (v: V) => void,
	resolve: (args: T) => void
}) {
	return (resolveArgs: T) => {
		return {
			...args.fields,
			create: args.create
		}
	}
}

class User {
	public id: string
	public name: string

	static get(id: string) {
		return {
			id,
			name: 'John'
		}
	}

	static getChatroomMembers(chatroomId: string): User[] {
		return []
	}
}

class Chatroom {
	public id: string

	static get(id: string) {
		return new Chatroom()
	}

	members() {
		return User.getChatroomMembers(this.id)
	}
}




useGraph(Chatroom.get("1").members())


// const user = GraphNode({
// 	resolve: (userId: number) => {
// 		return {
// 			id: userId,
// 			name: 'John'
// 		}
// 	},
// 	fields: {}
// })

// const chatroom = {
// 	create: (args: {
// 		name: string
// 	}) => {

// 	},
// 	get: (chatroomId: number) => {
// 		return {
// 			members: []
// 		}
// 	},
// 	members: (chatroomId: number) => {

// 	}
// }



// export class StringType extends GraphNode {
// 	public constructor() {
// 		super({
// 			fields: {}
// 		});
// 	}
// }

// export class ArrayType extends GraphNode {
// 	public constructor(type: GraphNode) {
// 		super({
// 			fields: {}
// 		});
// 	}
// }

// const user = new GraphNode({
// 	fields: {
// 		id: StringType,
// 		name: StringType
// 	}
// })

// const charoom = new GraphNode({
// 	fields: {
// 		id: StringType,
// 		members: new ArrayType(user)
// 	}
// });


// chatroom.