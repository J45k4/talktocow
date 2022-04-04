import React from "react"
import { Diary } from "../../src/components/diary/diary"
import { NavigationBar } from "../../src/components/navigation_bar"



export default function DiaryPage() {
    

    return (
        <div>
            <NavigationBar />
            <Diary /> 
        </div>
    )
}