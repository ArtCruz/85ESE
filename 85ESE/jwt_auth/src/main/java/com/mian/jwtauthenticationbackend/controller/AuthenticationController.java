package com.mian.jwtauthenticationbackend.controller;

import com.mian.jwtauthenticationbackend.model.AuthenticationResponse;
import com.mian.jwtauthenticationbackend.model.User;
import com.mian.jwtauthenticationbackend.service.AuthenticationService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@Tag(name = "Autenticação", description = "API para autenticação e registro de usuários")
@RestController
public class AuthenticationController {

    private final AuthenticationService authSerive;

    public AuthenticationController(AuthenticationService authSerive) {
        this.authSerive = authSerive;
    }

    @Operation(
        summary = "Registrar novo usuário",
        description = "Cria um novo usuário e retorna um token JWT.",
        responses = {
            @ApiResponse(responseCode = "200", description = "Usuário registrado com sucesso",
                content = @Content(schema = @Schema(implementation = AuthenticationResponse.class))),
            @ApiResponse(responseCode = "400", description = "Erro na requisição", content = @Content)
        }
    )
    @PostMapping("/register")
    public ResponseEntity<AuthenticationResponse> register(@RequestBody User request){
        return ResponseEntity.ok(authSerive.register(request));
    }

    @Operation(
        summary = "Autenticar usuário",
        description = "Autentica um usuário e retorna um token JWT.",
        responses = {
            @ApiResponse(responseCode = "200", description = "Usuário autenticado com sucesso",
                content = @Content(schema = @Schema(implementation = AuthenticationResponse.class))),
            @ApiResponse(responseCode = "400", description = "Credenciais inválidas", content = @Content)
        }
    )
    @PostMapping("/auth")
    public ResponseEntity<AuthenticationResponse> login(@RequestBody User request) {
        AuthenticationResponse response = authSerive.authenticationResponse(request);
        return ResponseEntity.ok(response);
    }

    @Operation(
        summary = "Testar API",
        description = "Endpoint de teste para verificar se a API está online.",
        responses = {
            @ApiResponse(responseCode = "200", description = "API funcionando", content = @Content)
        }
    )
    @GetMapping("/test")
    public String test() {
        return "API funcionando! 06/06-21:37";
    }
}