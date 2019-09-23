exports.divide = (req,res) =>{
    const {num1,num2} = req.body
    return res.status(200).send({result:num1/num2})
}

exports.multiply = (req,res) =>{
    const {num1,num2} = req.body
    return res.status(200).send({result:num1*num2})
}
exports.substract = (req,res) =>{
    const {num1,num2} = req.body
    return res.status(200).send({result:num1-num2})
}
exports.sum = (req,res) =>{
    const {num1,num2} = req.body
    return res.status(200).send({result:num1+num2})
}