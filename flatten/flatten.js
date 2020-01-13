function flattenArray(arrays) {
    let flat = [];
    flatten(flat, arrays);
    return flat
}


function flatten(flat, arrays) {

    for (let el of arrays) {
        if (Array.isArray(el)) {
            flatten(flat, el);
        } else {
            flat.push(el);
        }
    }

}

module.exports = flattenArray