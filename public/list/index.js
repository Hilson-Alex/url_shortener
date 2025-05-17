
const getData = () => {
    fetch('/short/list')
        .then(res => res.json())
        .then(populateTable)
}

const populateTable = (res) => {
    const table = document.getElementById('url-table').getElementsByTagName('tbody')[0]
    const cellOrder = ['key', 'originalUrl', 'shortUrl', 'expireDate']
    res.data.map(entry => {
        entry.expireDate = new Date(entry.expireDate * 1000).toLocaleDateString()
        return entry
    }).forEach(entry => {
        const row = table.insertRow()
        for (const key of cellOrder) {
            const cell = row.insertCell()
            cell.innerText = entry[key]
        }
    });
}