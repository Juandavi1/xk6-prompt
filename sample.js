import prompt from "k6/x/prompt"

export default function () {

    // Define options
    const options = ["smoke", "load"]

    // Read type from __ENV or Prompt
    const selected = __ENV.type ? __ENV.type : prompt.select("type of test:", ...options)

    // Print value
    console.log(selected)

}
