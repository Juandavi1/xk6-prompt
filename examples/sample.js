import prompt from "k6/x/prompt"
import http from 'k6/http'
import {check} from 'k6';


export const options = {
    vus: __ENV.vus ? __ENV.vus : prompt.readInt("total vus"),
    duration: __ENV.duration ? __ENV.duration : prompt.readString("duration seconds (s)"),
}

export default function () {

    // Define options
    const options = ["smoke", "load"]

    // Read type from __ENV or Prompt
    const selected = __ENV.type ? __ENV.type : prompt.select("type of test", ...options)
    const id = __ENV.env ? __ENV.env : prompt.readString("photo id (from 1 up to 100): ")

    // Print values
    console.log("values entered by the user: ", selected, selected3)

    let res = http.get('https://jsonplaceholder.typicode.com/photos/' + id)
    check(res, {
        'is status 200': (r) => r.status === 200,
    });


}
