<template>
  <div class="topnav">
    <a class="active" href="#home">Home</a>
    <input type="text" placeholder="Search..">
  </div>

  <p>Here is the contest</p>
  <table>
    <td>
      <tr>row1</tr>
      <tr>row2</tr>
      <tr>row3</tr>
    </td>
    <td>
      <tr>row1</tr>
      <tr>row2</tr>
      <tr>row3</tr>
    </td>
  </table>
</template>

<script>
export default {
  name: 'MainPage',
  props: {
    msg: String
  }
}
</script>

<style scoped>

</style>