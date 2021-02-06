
export const postJson = async (path: string, payload: any) => {
    const token = localStorage.getItem("token")

    const headers = {
        ["Content-Type"]: "application/json"
    }

    if (token) {
        headers["Authorization"] = "Bearer " + token
    }

    const res = await fetch(path, {
        method: "POST",
        headers: headers,
        body: JSON.stringify(payload)
    }).then(async (r) => {
        // if (r.status === 401) {
        //     window.localStorage.removeItem("token");
        //     window.location.href = "/";
        // }

        if (r.status !== 200) {
            const text = await r.text();

            throw new Error(text)
        }

        return r.json()
    });

    return res
}

export const getJson = async (path: string) => {
    const token = localStorage.getItem("token")

    const headers = {
        ["Content-Type"]: "application/json"
    }

    if (token) {
        headers["Authorization"] = "Bearer " + token
    }

    const res = await fetch(path, {
        headers: headers,
    }).then(r => {
        // if (r.status === 401) {
        //     window.localStorage.removeItem("token");
        //     window.location.href = "/";
        // }        

        return r.json()
    });

    return res
}