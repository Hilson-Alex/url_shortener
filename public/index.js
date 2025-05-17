
const shortUrl = (e) => {
    e.preventDefault()
    const data = new FormData(e.target)
    fetch('/short/create', {
        method: 'POST',
        body: JSON.stringify({
            originalUrl: data.get('originalUrl'),
            expireDate: Number(data.get('expireDate'))
        })
    }).then(res => res.json()).then(setResult)
}

const setResult =(data) => {
    document.getElementById('result').innerHTML = `
        <p>
            The URL ${data.originalUrl} was shortened to <a target="_blank" href=${data.shortUrl}>${data.shortUrl}</a>
        </p>
        <p onmouseover="showpop('expireInfo')" onmouseout="hidepop('expireInfo')">
            This link will be active until ${new Date(data.expireDate * 1000).toLocaleDateString()} ℹ️
        </p>
        <div id="expireInfo" class="tooltip" popover>The link may be active after this date, but don't take that as granted</div>
    `
}

const showpop = (id) => {
    console.log('aaa')
    document.getElementById(id).showPopover()
}

const hidepop = (id) => {
    document.getElementById(id).hidePopover()
}