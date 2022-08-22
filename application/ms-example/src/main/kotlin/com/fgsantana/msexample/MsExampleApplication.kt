package com.fgsantana.msexample

import co.elastic.apm.attach.ElasticApmAttacher
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class MsExampleApplication

fun main(args: Array<String>) {
	ElasticApmAttacher.attach()
	runApplication<MsExampleApplication>(*args)
}
