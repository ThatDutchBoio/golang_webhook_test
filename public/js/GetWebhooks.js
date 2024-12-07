window.onload = async () => {
    let webhookTable = document.getElementById('webhookTable');

    const generateElement = (id, host) => `<tr><td>${id}</td><td>${host}</td></tr>`

    const webHooks = await fetch('/api/getwebhooks', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })

    const body = await webHooks.json();

    console.log(body);

    for (let i = 0; i < body.length; i++) {
        webhookTable.innerHTML += generateElement(body[i].ID, body[i].Host)
    }

}