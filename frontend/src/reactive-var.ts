import { Observer, Subject, Subscription } from "rxjs";

export class ReactiveVar<T> {
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