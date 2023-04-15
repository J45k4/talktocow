import React, { useEffect, useState } from "react"

export const Window = (props: {
	title?: string
	show?: boolean
	x?: number
	y?: number
	width?: number
	height?: number
	children?: any
	onClose?: () => void
}) => {
	const [x, setX] = useState(100)
	const [y, setY] = useState(100)
	const [width, setWidth] = useState(200)
	const [height, setHeight] = useState(200)
	const [originalWidth, setOriginalWidth] = useState(200)
	const [originalHeight, setOriginalHeight] = useState(200)
	const [originalX, setOriginalX] = useState(0)
	const [originalY, setOriginalY] = useState(0)
	const [relativeX, setRelativeX] = useState(0)
	const [relativeY, setRelativeY] = useState(0)
	const [canResizeLeft, setCanResizeLeft] = useState(false)
	const [resizingLeft, setResizingLeft] = useState(false)
	const [canResizeRight, setCanResizeRight] = useState(false)
	const [resizingRight, setResizingRight] = useState(false)
	const [canResizeTop, setCanResizeTop] = useState(false)
	const [resizingTop, setResizingTop] = useState(false)
	const [canResizeBottom, setCanResizeBottom] = useState(false)
	const [resizingBottom, setResizingBottom] = useState(false)

	const [moving, setMoving] = useState(false)
	const [canMove, setCanMove] = useState(false)

	// const [shown, setShown] = useState(true)

	useEffect(() => {
		if (!props.x) {
			return
		}

		setX(props.x)
	}, [props.x, setX])

	useEffect(() => {
		if (!props.y) {
			return
		}

		setY(props.y)
	}, [props.y, setY])

	useEffect(() => {
		if (!props.width) {
			return
		}

		setWidth(props.width)
	}, [props.width, setWidth])

	useEffect(() => {
		if (!props.height) {
			return
		}

		setHeight(props.height)
	}, [props.height, setHeight])


	useEffect(() => {
		function onMouseMove(e: MouseEvent) {
			const clientX = e.clientX
			const clientY = e.clientY

			if (resizingLeft) {
				setWidth(originalWidth + originalX - clientX)
				setX(clientX)
			}

			if (resizingRight) {
				setWidth(originalWidth + clientX - originalX)
			}

			if (resizingTop) {
				setHeight(originalHeight + originalY - clientY)
				setY(clientY)
			}

			if (resizingBottom) {
				setHeight(originalHeight + clientY - originalY)
			}

			if (moving) {
				setX(e.clientX - relativeX)
				setY(e.clientY - relativeY)
			}

			if (clientX > x - 5 && clientX < x + 5) {
				setCanResizeLeft(true)
			} else {
				setCanResizeLeft(false)
			}

			if (clientX > x + width - 5 && clientX < x + width + 5) {
				setCanResizeRight(true)
			} else {
				setCanResizeRight(false)
			}

			if (clientY > y - 5 && clientY < y + 5) {
				setCanResizeTop(true)
			} else {
				setCanResizeTop(false)
			}

			if (clientY > y + height - 5 && clientY < y + height + 5) {
				setCanResizeBottom(true)
			} else {
				setCanResizeBottom(false)
			}

			if (clientX >= x && clientX <= x + width && clientY >= y && clientY <= y + 30) {
				setCanMove(true)
			} else {
				setCanMove(false)
			}
		}

		

		function onMouseDown(e: MouseEvent) {
			if (canResizeLeft || canResizeRight || canResizeTop || canResizeBottom) {
				e.preventDefault()

				setResizingLeft(canResizeLeft)
				setResizingRight(canResizeRight)
				setResizingTop(canResizeTop)
				setResizingBottom(canResizeBottom)

				setOriginalX(e.clientX)
				setOriginalY(e.clientY)
				setOriginalWidth(width)
				setOriginalHeight(height)
			} else if (canMove) {
				e.preventDefault()

				setMoving(canMove)
				setRelativeX(e.clientX - x)
				setRelativeY(e.clientY - y)
			}
		}

		

		function onmouseup(e: MouseEvent) {
			setMoving(false)
			setResizingLeft(false)
			setResizingRight(false)
			setResizingTop(false)
			setResizingBottom(false)
		}

		window.addEventListener("mousemove", onMouseMove)
		window.addEventListener("mousedown", onMouseDown)
		window.addEventListener("mouseup", onmouseup)

		return () => {
			window.removeEventListener("mousemove", onMouseMove)
			window.removeEventListener("mousedown", onMouseDown)
			window.removeEventListener("mouseup", onmouseup)
		}
	}, [
		moving, 
		canResizeLeft, 
		canResizeRight, 
		canResizeTop, 
		canResizeBottom, 
		resizingLeft,
		resizingRight, 
		resizingTop, 
		resizingBottom, 
		canMove, 
		x, 
		width,
		y, 
		height, 
		originalWidth, 
		originalX, 
		originalHeight, 
		originalY, 
		relativeX, 
		relativeY
	])

	// if (!shown) {
	// 	return <></>
	// }
	
	let cursor

	if (canResizeLeft && canResizeTop) {
		cursor = "nw-resize"
	} else if (canResizeRight && canResizeBottom) {
		cursor = "nw-resize"
	} else if (canResizeLeft || canResizeRight) {
		cursor = "ew-resize"
	} else if (canResizeTop || canResizeBottom) {
		cursor = "ns-resize"
	} else if (canMove) {
		cursor = "move"
	}

	return (
		<div style={{
			position: "absolute",
			left: `${x}px`,
			top: `${y}px`,
			width: `${width}px`,
			height: `${height}px`,
			border: "1px solid black",
			backgroundColor: "white",
			cursor: cursor,
			display: "flex",
			flexDirection: "column",
		}}>
			<div style={{
				backgroundColor: "#5F5E5E",
				width: "100%",
				height: "30px",
				overflow: "hidden"
			}}>
				<div style={{
					display: "flex"
				}}>
					<div style={{
						flexGrow: 1
					}}>
						{props.title}
					</div>
					<div>
						<button onClick={() => {
							// setShown(false)
							if (props.onClose) {
								props.onClose()
							}
						}}>
							close
						</button>
					</div>
				</div>
			</div>
			<div style={{
				backgroundColor: "white",
				flexGrow: 1,
				overflow: "auto"
			}}>
				{props.children}
			</div>
		</div>
	)
}