package com.mian.jwtauthenticationbackend.controller;

import com.mian.jwtauthenticationbackend.model.AuthenticationResponse;
import com.mian.jwtauthenticationbackend.model.User;
import com.mian.jwtauthenticationbackend.service.AuthenticationService;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class AuthenticationController {

    private final AuthenticationService authSerive;

    public AuthenticationController(AuthenticationService authSerive) {
        this.authSerive = authSerive;
    }
    @PostMapping("/register")
    public ResponseEntity<AuthenticationResponse> register(
            @RequestBody User request){
        return ResponseEntity.ok(authSerive.register(request));
    }
    // @PostMapping("/login")
    // public ResponseEntity<AuthenticationResponse> login(
    //         @RequestBody User request)
    // {
    //     return ResponseEntity.ok(authSerive.authenticationResponse(request));
    // }

    @PostMapping("/auth")
    public ResponseEntity<AuthenticationResponse> login(
            @RequestBody User request)
    {
        return ResponseEntity.ok(authSerive.authenticationResponse(request));
    }

    @GetMapping("/test")
    public String test() {
        return "API funcionando! 06/06-21:37";
    }


}