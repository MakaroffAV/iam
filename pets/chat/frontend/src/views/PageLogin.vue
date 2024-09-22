<template>
    <BlankTemplate>
        <template v-slot:content>
            <div class="signin-wrapper">
                <div class="signin-box">
                    <h3 class="signin-title-secondary">
                        {{ this.signInMessage }}
                    </h3>
                    <FormLogin />
                </div>
            </div>
        </template>
    </BlankTemplate>
</template>

<script>

import FormLogin from "../module/FormLogin.vue"

export default {

    components: {
        FormLogin,
    },

    created() {
        document.title  = "Войти"
    },

    data() {
        return {
            signUpMessage: "Нет аккаунта?",
            signInMessage: "Войдите для доступа к платформе"
        }
    }

}
</script>
