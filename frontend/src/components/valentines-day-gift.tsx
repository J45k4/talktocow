import { useEffect, useState } from "react";
import styles from "./valentines-day-gift.module.css";

export const ValentinesDayGift = (props: {
    onContinue: () => void
}) => {
    const [firstVisible, setFirstVisible] = useState(true)

    // useEffect(() => {
    //     setTimeout(() => {
    //         setFirstVisible(false)
    //     }, 3000)
        
    // })

    return (
        // <div style={{
        //     position: "absolute",
        //     top: "-50px"
        // }}>
            <div className={styles.rootContainer}>
                {firstVisible &&
                <div className={styles.firstGroup}> 
                    Good morning 😘
                </div>}
                <div className={styles.secondGroup}>
                    Love you 😍
                </div>
                <div className={styles.thirthGroup}>
                    I am lucky to have you 🥰
                </div>
                <div className={styles.fourthGroup}>
                    Happy valentine's day my love 🥰🥰
                </div>
                <div style={{
                    display: "flex",
                    justifyContent: "center"
                }}>
                <button className={styles.continueButton} onClick={props.onContinue}>
                    Continue
                </button>
                </div>                
            </div>
        // </div>
    )
}