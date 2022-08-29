package com.fgsantana.msktexample

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/kt")
class ExController {

    @GetMapping
    fun get()  = "GET test"

    @PostMapping
    fun post()  = "POST test"

}