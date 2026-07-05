import React, { useEffect, useRef, useState } from "react"
import { FaMicrophone, FaStop } from "react-icons/fa"
import { useNavigate } from "react-router-dom"
import { postFormData } from "../../api-methods"
import { Modal } from "../modal"
import styles from "./diary.module.css"

type UploadedRecording = {
    id: number
    fileName: string
}

const maxRecordingDurationMs = 30 * 60 * 1000

const recordingMimeTypeCandidates = [
    "audio/webm;codecs=opus",
    "audio/webm",
    "audio/ogg;codecs=opus",
    "audio/ogg",
    "audio/mp4"
]

const formatElapsed = (elapsedMs: number) => {
    const totalSeconds = Math.floor(elapsedMs / 1000)
    const minutes = Math.floor(totalSeconds / 60)
    const seconds = totalSeconds % 60

    return `${minutes}:${seconds.toString().padStart(2, "0")}`
}

const supportedRecordingMimeType = () => {
    if (typeof MediaRecorder === "undefined") {
        return ""
    }

    return recordingMimeTypeCandidates.find(type => MediaRecorder.isTypeSupported(type)) ?? ""
}

const recordingExtension = (mimeType: string) => {
    if (mimeType.includes("ogg")) {
        return "ogg"
    }

    if (mimeType.includes("mp4")) {
        return "m4a"
    }

    return "webm"
}

export function DiaryRecorder() {
    const navigate = useNavigate()
    const [elapsedMs, setElapsedMs] = useState(0)
    const [error, setError] = useState("")
    const [isOpen, setIsOpen] = useState(false)
    const [status, setStatus] = useState<"idle" | "requesting" | "recording" | "uploading">("idle")
    const chunksRef = useRef<Blob[]>([])
    const intervalRef = useRef<number | null>(null)
    const maxTimeoutRef = useRef<number | null>(null)
    const recorderRef = useRef<MediaRecorder | null>(null)
    const shouldUploadRef = useRef(false)
    const startTimeRef = useRef(0)
    const streamRef = useRef<MediaStream | null>(null)

    const clearTimers = () => {
        if (intervalRef.current != null) {
            window.clearInterval(intervalRef.current)
            intervalRef.current = null
        }

        if (maxTimeoutRef.current != null) {
            window.clearTimeout(maxTimeoutRef.current)
            maxTimeoutRef.current = null
        }
    }

    const stopStream = () => {
        streamRef.current?.getTracks().forEach(track => track.stop())
        streamRef.current = null
    }

    const resetRecorder = () => {
        clearTimers()
        stopStream()
        recorderRef.current = null
        chunksRef.current = []
        shouldUploadRef.current = false
    }

    useEffect(() => {
        return () => resetRecorder()
    }, [])

    const uploadRecording = async (recorder: MediaRecorder) => {
        const durationMs = Date.now() - startTimeRef.current
        const mimeType = recorder.mimeType || supportedRecordingMimeType() || "audio/webm"
        const blob = new Blob(chunksRef.current, { type: mimeType })
        const fileName = `diary-recording-${new Date().toISOString().replace(/[:.]/g, "-")}.${recordingExtension(mimeType)}`
        const formData = new FormData()
        formData.append("recording", new File([blob], fileName, { type: mimeType }))

        const response = await postFormData<UploadedRecording>("/api/diary/recording", formData)

        if (response.error) {
            throw new Error(response.error.message)
        }

        if (!response.payload) {
            throw new Error("Recording upload did not return a file")
        }

        const params = new URLSearchParams({
            recordingDurationMs: String(durationMs),
            recordingFileId: String(response.payload.id),
            recordingFileName: response.payload.fileName
        })

        navigate(`/diary/new?${params.toString()}`)
    }

    const finishRecording = () => {
        const recorder = recorderRef.current

        if (!recorder || recorder.state === "inactive") {
            return
        }

        shouldUploadRef.current = true
        setStatus("uploading")
        clearTimers()
        recorder.stop()
    }

    const cancelRecording = () => {
        shouldUploadRef.current = false

        if (recorderRef.current && recorderRef.current.state !== "inactive") {
            recorderRef.current.stop()
        }

        resetRecorder()
        setElapsedMs(0)
        setError("")
        setIsOpen(false)
        setStatus("idle")
    }

    const startRecording = async () => {
        setIsOpen(true)
        setError("")
        setElapsedMs(0)

        if (!navigator.mediaDevices?.getUserMedia || typeof MediaRecorder === "undefined") {
            setError("Recording is not supported by this browser.")
            setStatus("idle")
            return
        }

        const mimeType = supportedRecordingMimeType()

        if (!mimeType) {
            setError("This browser cannot record a supported audio format.")
            setStatus("idle")
            return
        }

        setStatus("requesting")

        try {
            const stream = await navigator.mediaDevices.getUserMedia({ audio: true })
            const recorder = new MediaRecorder(stream, {
                audioBitsPerSecond: 32000,
                mimeType
            })

            streamRef.current = stream
            recorderRef.current = recorder
            chunksRef.current = []
            shouldUploadRef.current = false

            recorder.ondataavailable = event => {
                if (event.data.size > 0) {
                    chunksRef.current.push(event.data)
                }
            }

            recorder.onstop = () => {
                clearTimers()
                stopStream()

                if (!shouldUploadRef.current) {
                    return
                }

                void uploadRecording(recorder).catch(uploadError => {
                    setError(uploadError instanceof Error ? uploadError.message : String(uploadError))
                    setStatus("idle")
                })
            }

            startTimeRef.current = Date.now()
            recorder.start(1000)
            setStatus("recording")
            intervalRef.current = window.setInterval(() => setElapsedMs(Date.now() - startTimeRef.current), 250)
            maxTimeoutRef.current = window.setTimeout(finishRecording, maxRecordingDurationMs)
        } catch (recordingError) {
            resetRecorder()
            setError(recordingError instanceof Error ? recordingError.message : String(recordingError))
            setStatus("idle")
        }
    }

    return (
        <>
            <button className={styles.recordEntryButton} onClick={startRecording} type="button">
                <FaMicrophone />
                Record
            </button>
            <Modal isOpen={isOpen} title="Record diary" onClose={() => {
                if (status !== "uploading") {
                    cancelRecording()
                }
            }}>
                <div className={styles.recordingModal}>
                    <div className={styles.recordingTimer}>{formatElapsed(elapsedMs)}</div>
                    {status === "requesting" && <div className={styles.recordingStatus}>Waiting for microphone...</div>}
                    {status === "recording" && <div className={styles.recordingStatus}>Recording</div>}
                    {status === "uploading" && <div className={styles.recordingStatus}>Uploading recording...</div>}
                    {error && <div className={styles.saveError}>{error}</div>}
                    <div className={styles.recordingActions}>
                        <button className={styles.secondaryButton} onClick={cancelRecording} disabled={status === "uploading"} type="button">
                            Cancel
                        </button>
                        <button className={styles.primaryButton} onClick={finishRecording} disabled={status !== "recording"} type="button">
                            <FaStop />
                            Stop
                        </button>
                    </div>
                </div>
            </Modal>
        </>
    )
}
