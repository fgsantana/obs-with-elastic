package com.fgsantana.msexample

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class ExController {

    @GetMapping
    fun get()  = "GET test"

    @PostMapping
    fun post()  = "POST test"

}