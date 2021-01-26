import React, { ReactFragment } from "react";

export const IsLoggedIn = (props: {
    children: any
}) => {
    if (typeof window == "undefined") {
        return <React.Fragment></React.Fragment>
    }

    const token = localStorage.getItem("token");

    if (token) {
        return <React.Fragment>
            {props.children}
        </React.Fragment>
    }

    return <React.Fragment />
}

export const IsNotLoggedIn = (props: {
    children: any
}) => {
    if (typeof window == "undefined") {
        return <React.Fragment></React.Fragment>
    }

    const token = localStorage.getItem("token");

    if (token) {
        return <React.Fragment />
    }


    return <React.Fragment>
        {props.children}
    </React.Fragment>
}