plugins {
    application
    id("com.gradleup.shadow") version "9.3.0"
}

repositories {
    mavenCentral()
}

dependencies {
    implementation(libs.guava)
    implementation("org.slf4j:slf4j-simple:2.0.9")
    implementation ("org.java-websocket:Java-WebSocket:1.5.4")
}

testing {
    suites {
        val test by getting(JvmTestSuite::class) {
            useJUnit("4.13.2")
        }
    }
}


java {
    toolchain {
        languageVersion = JavaLanguageVersion.of(25)
    }
}

application {
    mainClass = "org.ws_server.Server"
}
