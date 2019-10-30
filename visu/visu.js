const main = document.querySelector('#main')
const maxSpeed = 10
let stepIndex = 0
let speed = 5

function setupGrid() {
    main.innerHTML = ''
    current.forEach((row, rowIndex) => {
        const newDiv = document.createElement('div')
        newDiv.classList.add('row')
        row.forEach((nbr, colIndex) => {
            const div = document.createElement('div')
            if (nbr != 0) {
                if (nbr == wanted[rowIndex][colIndex]) {
                    div.classList.add('well-placed')
                }
                const p = document.createElement('p')
                p.innerHTML = nbr
                div.appendChild(p)
            } else {
                div.classList.add('empty')
            }
            newDiv.appendChild(div)
        })
        main.append(newDiv)
    })
}
/**
 * @param {Number} nbr
 * @returns {Element}
 */
function getNumberDiv(nbr) {
    let row, col
    current.forEach((r, rowIndex) => {
        r.forEach((n, colIndex) => {
            if (nbr == n) {
                row = rowIndex
                col = colIndex
            }
        })
    })

    if (row == undefined && col == undefined) {
        return null
    }

    return main.children.item(row).children.item(col)
}
/**
 * @param {Element} div
 * @param {String} direction
 * @returns {Element}
 */
function getSideDiv(div, direction) {
    switch (direction) {
        case 'up':
            return div.parentElement.previousElementSibling.children.item(Array.prototype.indexOf.call(div.parentElement.children, div))
        case 'down':
            return div.parentElement.nextElementSibling.children.item(Array.prototype.indexOf.call(div.parentElement.children, div))
        case 'left':
            return div.previousElementSibling
        case 'right':
            return div.nextElementSibling
    }
    return null
}
/**
 * 
 * @param {Number} nbr 
 * @param {String} direction 
 */
function moveNumberInGrid(nbr, direction) {
    let i, j
    current.forEach((row, rowIndex) => {
        row.forEach((nb, colIndex) => {
            if (nbr == nb) {
                i = rowIndex
                j = colIndex
            }
        })
    })

    let tmp = current[i][j]
    switch (direction) {
        case 'up':
            current[i][j] = current[i - 1][j]
            current[i - 1][j] = tmp
            return
        case 'down':
            current[i][j] = current[i + 1][j]
            current[i + 1][j] = tmp
            return
        case 'left':
            current[i][j] = current[i][j - 1]
            current[i][j - 1] = tmp
            return
        case 'right':
            current[i][j] = current[i][j + 1]
            current[i][j + 1] = tmp
            return
    }
}

function prepareDivs(...divs) {
    divs.forEach(div => {
        div.classList.remove(...div.classList)
        div.innerHTML = ''
    })
}

function moveNumber(nbr, direction) {
    const nbrDiv = getNumberDiv(nbr)
    const sideDiv = getSideDiv(nbrDiv, direction)
    const margin = parseInt(window.getComputedStyle(nbrDiv).margin)
    const width = parseInt(window.getComputedStyle(nbrDiv).width)
    const height = parseInt(window.getComputedStyle(nbrDiv).height)
    let incr = 0
    let intervalId
    switch (direction) {
        case 'up':
            intervalId = setInterval(() => {
                if (incr == 0) {
                    prepareDivs(nbrDiv, sideDiv)
                    nbrDiv.style.marginTop = '0px'
                    sideDiv.style.marginBottom = '0px'
                }
                incr++
                nbrDiv.style.marginBottom = (margin + incr < width + 2 * margin ? margin + incr : width + 2 * margin) + 'px'
                nbrDiv.style.height = (height - incr) + 'px'
                sideDiv.style.marginTop = (margin + height - incr > margin ? margin + height - incr : margin) + 'px'
                sideDiv.style.height = (margin + incr < height ? margin + incr : height) + 'px'
            }, speed)
            break
        case 'down':
            intervalId = setInterval(() => {
                if (incr == 0) {
                    prepareDivs(nbrDiv, sideDiv)
                    nbrDiv.style.marginBottom = '0px'
                    sideDiv.style.marginTop = '0px'
                }
                incr++
                nbrDiv.style.marginTop = (2 * margin + incr < width + 2 * margin ? 2 * margin + incr : width + 2 * margin) + 'px'
                nbrDiv.style.height = (height - incr) + 'px'
                sideDiv.style.marginBottom = (margin + height - incr > margin ? margin + height - incr : margin) + 'px'
                sideDiv.style.height = (margin + incr < height ? margin + incr : height) + 'px'
            }, speed)
            break
        case 'left':
            intervalId = setInterval(() => {
                if (incr == 0) {
                    prepareDivs(nbrDiv, sideDiv)
                    nbrDiv.style.marginLeft = '0px'
                    sideDiv.style.marginRight = '0px'
                }
                incr++
                nbrDiv.style.marginRight = (margin + incr) + 'px'
                nbrDiv.style.width = (2 * margin + width - incr) + 'px'
                sideDiv.style.marginLeft = (width - incr > margin ? width - incr : margin) + 'px'
                sideDiv.style.width = (margin + incr < width ? margin + incr : width) + 'px'
            }, speed)
            break
        case 'right':
            intervalId = setInterval(() => {
                if (incr == 0) {
                    prepareDivs(nbrDiv, sideDiv)                    
                    nbrDiv.style.marginRight = '0px'
                    sideDiv.style.marginLeft = '0px'
                }
                incr++
                nbrDiv.style.marginLeft = (margin + incr) + 'px'
                nbrDiv.style.width = (2 * margin + width - incr) + 'px'
                sideDiv.style.marginRight = (width - incr > margin ? width - incr : margin) + 'px'
                sideDiv.style.width = (margin + incr < width ? margin + incr : width) + 'px'
            }, speed)
            break
    }

    setTimeout(() => {
        if (intervalId) {
            clearInterval(intervalId)
        }
        sideDiv.innerHTML = `<p>${nbr}</p>`
        moveNumberInGrid(nbr, direction)
        setupGrid()

        // Launching next step
        stepIndex++
        if (stepIndex < steps.length) {
            moveNumber(steps[stepIndex].nbr, steps[stepIndex].dir)
        }
    }, speed * 110)
}

setupGrid()

document.querySelector('#start').addEventListener('click', function(e) {
    e.preventDefault()
    this.parentElement.remove()
    moveNumber(steps[stepIndex].nbr, steps[stepIndex].dir)
})

document.querySelector('#speed').addEventListener('change', function () {
    speed = maxSpeed - parseInt(this.value) + 1
})
