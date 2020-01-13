var Flatten = require("../flatten")

describe("Flatten", function() {
    it("flattens nested array", function() {
        let flat = Flatten([1, [2, [3]], 4]);
        expect(flat).toEqual([ 1, 2, 3, 4 ]);
    });
});