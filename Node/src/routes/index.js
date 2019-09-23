const {Router} =require('express')
const router = Router()
const {divide,substract,sum,multiply} = require('../controllers')


router.get('/',(req,res)=>res.json({message:"node in docker"}))
router.post('/divide',divide)
router.post('/multiply',multiply)
router.post('/sum',sum)
router.post('/substract',substract)

module.exports = router