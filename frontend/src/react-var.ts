import { Observer, Subject, Subscription } from "rxjs";

export class ReactVar<T> {
	private state: T
	private subject: Subject<T>
	
	public constructor(initialState: T) {
		this.state = initialState
		this.subject = new Subject()
	}

	public set(newState: T) {
		this.state = newState
		this.subject.next(newState)
	}

	public get(): T {
		return this.state
	}

	public sub(observerOrNext?: Partial<Observer<T>> | ((value: T) => void)): Subscription {
		return this.subject.subscribe(observerOrNext)
	}
}


export class ReactVars<K, T> {
	private map = new Map<K, ReactVar<T>>()

	public get(key: K): ReactVar<T> {
		let rv = this.map.get(key)

		if (!rv) {
			rv = new ReactVar(null)
			this.map.set(key, rv)
		}

		return rv
	}
}