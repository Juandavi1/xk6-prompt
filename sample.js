import prompt from "k6/x/prompt"

export default function () {

    // Define options
    const options = ["smoke", "load"]

    // Read type from __ENV or Prompt
    const selected = __ENV.type ? __ENV.type : prompt.select("type of test", ...options)
    const selected2 = __ENV.type ? __ENV.VUS : prompt.readInt("total vus")
    const selected3 = __ENV.type ? __ENV.ENV : prompt.readString("environment")

    // Print value
    console.log(selected, selected2, selected3)

}
