package com.fgsantana.msktexample

import co.elastic.apm.attach.ElasticApmAttacher
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class MsKtExampleApplication

fun main(args: Array<String>) {
	ElasticApmAttacher.attach()
	runApplication<MsKtExampleApplication>(*args)
}
