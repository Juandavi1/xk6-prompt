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
    const selected3 = __ENV.env ? __ENV.env : prompt.readString("any text")

    // Print values
    console.log("values entered by the user: ", selected, selected3)

    let res = http.get('http://httpbin.org/cookies')
    check(res, {
        'is status 200': (r) => r.status === 200,
    });


}
