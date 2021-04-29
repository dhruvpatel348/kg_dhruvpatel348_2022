const args = process.argv.slice(2)

var a = ["Zero","One", "Two", "Three", "Four", "Five","Six","Seven","Eight","Nine"];

var foo = new Array(args.length);
var final = new Array(args.length)

for (let i = 0; i< args.length;i++) {
    foo[i] = Number(args[i])
    var val = foo[i],
    split = [];
    var poo = new Array(split.length)

    while (val) {
    split.push(val % 10);
    val = parseInt(val/10);
   }
   split.reverse()
   for(let j=0;j<split.length;j++)
   {
       poo[j] = a[split[j]]
   }
   final[i] = poo.join("")

}
console.log(final.join(","))