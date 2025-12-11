package main

type contextKey string

const isAuthenticatedKeyContextKey = contextKey("isAuthenticated")
const authenticatedUserContextKey = contextKey("authenticatedUser")
