import React, { useEffect } from "react"

import { getJson } from "../../../src/utility/talktocow-api-helpers"
import { useRouter } from "next/dist/client/router"

export default function DiaryEntryPage() {
    const router = useRouter()

    console.log("query", router.query)

    const diaryEntryId = router.query.diaryEntryId

    useEffect(() => {
        getJson("/api/diary/entry/" + diaryEntryId)
    }, [diaryEntryId])

    return (
        <div>
            <div>
                Title
                <input type="text" />
            </div>
            <div>
                Body
                <textarea />
            </div>
            <button>
                Update
            </button>
        </div>
    )
}