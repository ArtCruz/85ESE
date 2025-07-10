package com.mian.jwtauthenticationbackend.model;

import io.swagger.v3.oas.annotations.media.Schema;

@Schema(description = "Resposta de autenticação contendo o token JWT")
public class AuthenticationResponse {
    @Schema(description = "Token JWT gerado", example = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...")
    private String token;

    public AuthenticationResponse(String token){
        this.token = token;
    }
    public String getToken(){
        return token;
    }
}