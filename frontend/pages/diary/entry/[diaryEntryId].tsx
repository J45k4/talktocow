import React, { useEffect, useState } from "react"
import { getJson, putJson } from "../../../src/utility/talktocow-api-helpers"

import { useRouter } from "next/dist/client/router"

export default function DiaryEntryPage() {
    const router = useRouter()

    const diaryEntryId = router.query.diaryEntryId

    const [entry, setEntry] = useState<any>()

    useEffect(() => {
        if (!diaryEntryId) {
            return
        }

        getJson("/api/diary/entry/" + diaryEntryId).then(r => {
            setEntry(r.payload)
        })
    }, [diaryEntryId])

    const [newFiles, setNewFiles] = React.useState<File[]>([])
    const [images, setImages] = React.useState([])

    return (
        <div style={{
            display: "flex",
            flexDirection: "column",
            padding: "10px",
            height: "90%"
        }}>
            <div style={{

            }}>
                <button onClick={() => {
                    router.back()
                }}>
                    back
                </button>
                Title
                <input type="text" value={entry?.title} onChange={e => {
                    setEntry({
                        ...entry,
                        title: e.target.value
                    })
                }} />
            </div>
            <div style={{
                flexGrow: 1,
                display: "flex",
                flexDirection: "column",
            }}>
                <textarea value={entry?.body} style={{
                    flexGrow: 1
                }} onChange={e => {
                    setEntry({
                        ...entry,
                        body: e.target.value
                    })
                }} />                
            </div>
            <div>
                {images.map((f, i) => {
                    return <div key={i}><img src={f} style={{
                        width: "200px"
                    }} /></div>
                })}
            </div>
            <div style={{

            }}>
                <button onClick={() => {
                    putJson("/api/diary/entry/" + diaryEntryId, {
                        title: entry.title,
                        body: entry.body,
                        mask: ["title", "body"]
                    })
                }}>
                    Update
                </button>
                <input type="file"  multiple onChange={e => {
                    e.preventDefault()
                    console.log(e)
                    void (async () => {
                        const files = (e.target as any).files

                        const newImages = [];

                        for (const f of files) {
                            const result = await new Promise((resolve, reject) => {
                                const reader = new FileReader();
                                reader.onload = function (e) {
                                    resolve(e.target.result)
                                }
                                reader.readAsDataURL(f);
                            }) 

                            newImages.push(result)
                        }

                        setImages([
                            ...images,
                            ...newImages
                        ])

                        setNewFiles([...newFiles, ...files])
                    })()

                    
                }} />
            </div>
        </div>
    )
}