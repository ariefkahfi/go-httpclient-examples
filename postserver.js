const express = require('express')
const app = express()
const multer = require('multer')
const mime = require('mime')

app.use(express.urlencoded({extended:true}))


const dStorage = multer.diskStorage({
    destination:__dirname + '/uploads',
    filename:(req,file,cb)=>{
        cb(null , `${file.fieldname}`)
    }
})

const upload = multer({storage: dStorage})




app.post('/api/doUpload',(req,res,next)=>{
    console.log(req.get('content-type'),req.headers)
    console.log(req.body)
    next()
},upload.single('file_upload'),(req,res)=>{
    res.json({
        code:200,
        body:req.body,
        filePath:req.file.filename
    })
})

app.post('/api/doPost',(req,res)=>{
    console.log(`incoming client`,req.ip)
    if(!req.body.name  || !req.body.msg ) {
        res.status(400).json({
            code:'400',
            message:'BAD_REQUEST',
            data:'Invalid request body , required (name , msg) parameters'
        })
    }

    res.json({
        code:'200',
        message:"ECHO_RESPONSE",
        data:req.body
    })
})



app.listen(9600 , ()=> console.log('listening on port 9600'))
