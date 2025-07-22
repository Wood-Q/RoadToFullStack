<script>
  let username=$state('Woodq')
  $effect(()=>{
    console.log(`Username changed to ${username}`)
  })
</script>

<h1>Hello {username}</h1>
<input type="text" bind:value={username}/>