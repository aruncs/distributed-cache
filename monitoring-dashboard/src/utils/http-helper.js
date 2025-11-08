export async function makeHttpGet(url, headers) {
    let response = await fetch(url)
    let result = await response.json()
    return result
}