

function ShowTwoBtnText(Element) {
    Element.style.backgroundColor = 'white';
}

function setLowFreq() {
    el = getData();

    if (el.Freq > 1) {
        el.Freq -= 1
    }

    let freqq = el.Freq;
    let idd = el.Id;

    let response = fetch('http://localhost:5500/set', { method: 'post', body: JSON.stringify({ freqq, idd }) })
}

function setHiFreq() {
    el = getData();

    if (el.Freq < 3) {
        el.Freq += 1
    }

    let freqq = el.Freq;
    let idd = el.Id;

    let response = fetch('http://localhost:5500/set', { method: 'post', body: JSON.stringify({ freqq, idd }) })
}

function setFreqZero() {
    el = getData();

    el.Freq = 0;

    let freqq = el.Freq;
    let idd = el.Id;

    let response = fetch('http://localhost:5500/set', { method: 'post', body: JSON.stringify({ freqq, idd }) })
}

function getData() {
    let response = fetch('http://localhost:5500/get', { method: 'post', headers: { "Content-Type": "application/json" } })
        .then((resp) => resp.json())
        .then(function (data) {
            elm = data;
            randEl = Math.floor(Math.random() * 3)
            if (randEl == 1 && data.Freq == 1) {
                toHtml(data);
                document.getElementById('get2').removeAttribute('disabled');
                document.getElementById('get1').setAttribute('disabled', true);
            } else if (randEl > 1 && data.Freq == 2) {
                toHtml(data);
                document.getElementById('get2').removeAttribute('disabled');
            } else if (data.Freq == 3) {
                toHtml(data);
                document.getElementById('get2').setAttribute('disabled', true);
                document.getElementById('get1').removeAttribute('disabled');
            } else {
                getData()
            }
        })
        .catch(function (error) {
            console.log(error)
        })

    s = document.getElementById("twobtn").style;
    if (s.backgroundColor != 'black') {
        s.backgroundColor = 'black';
    }

    return elm
}

function toHtml(data) {
    document.getElementById('one').innerHTML = data.Fword;
    document.getElementById('answer').innerHTML = data.Sword;
}

function HiddenLearn(Element) {
    if (Element.style.opacity == 1) {
        Element.style.opacity = 0;
    } else {
        Element.style.opacity = 1;
    }
}

document.getElementById("get").addEventListener("click", getData);
document.getElementById("get0").addEventListener("click", getData);
document.getElementById("get1").addEventListener("click", setLowFreq);
document.getElementById("get2").addEventListener("click", setHiFreq);
document.getElementById("get3").addEventListener("click", setFreqZero);