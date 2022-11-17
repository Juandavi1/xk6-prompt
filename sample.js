import prompt from "k6/x/prompt"

export default function () {

    // Define options
    const options = ["smoke", "load"]

    // Read type from __ENV or Prompt
    const selected = __ENV.type ? __ENV.type : prompt.select("type of test:", ...options)
    const selected2 = __ENV.type ? __ENV.type : prompt.readInt("xd:")
    const selected3 = __ENV.type ? __ENV.type : prompt.readString("xd2")

    // Print value
    console.log(selected, selected2, selected3)

}
