<script lang="ts">
    let login = false;
    
    function submitLogin(e:MouseEvent) {
        e.preventDefault();
    }

    function submitRegister(e:MouseEvent) {
        e.preventDefault();
    }

</script>
<div id="center">
    {#if login}
        <div id="loginWrapper">
            <h1>Login</h1>
            <form on:submit={submitLogin} method="post" class="formCenter">
                <label>
                    Email:
                    <input required type="email" name="email" id="email">
                </label>
                <label>
                    Password:    
                    <input required type="password" name="password" id="password">
                </label>
                <input type="submit" value="Login">
            </form>
            <p class="switchRegisterLogin" on:click={e => login = !login}>Rather register</p>
        </div>
        
    {:else}
        <div id="registerWrapper">
            <h1>Register</h1>
            <form on:submit={submitRegister} method="post" class="formCenter">
                <label>
                    Email:
                    <input required type="email" name="email" id="email">
                </label>
                <label>
                    Password:    
                    <input required type="password" name="password" id="password">
                </label>
                <label>
                    Username:
                    <input required type="text" name="username" id="username">
                </label>
                <input type="submit" value="Register">
            </form>
        
        <p class="switchRegisterLogin" on:click={e => login = !login}>Rather login</p>
        </div>
    {/if}
</div>
<style lang="scss">

#center {
    width: 100%;
    min-height: 100vh;
    
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.formCenter {
    display: flex;
    flex-direction: column;
    
    input {
        display: block;
    }
    
}

.switchRegisterLogin {
cursor: pointer;
color: gray;
}
    
</style>