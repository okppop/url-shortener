const BASE_URL = window.location.origin

export async function CreateURL(data) {
    const response = await fetch(`${BASE_URL}/api/url`, {
        method: "POST",
        body: JSON.stringify(data),
        headers: {
            "content-type": "application/json"
        }
    })
    if (response.status !== 201) {
        throw response.status
    }
    const responseData = await response.json()
    return responseData
}