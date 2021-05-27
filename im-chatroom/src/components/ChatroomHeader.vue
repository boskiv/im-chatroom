<template>
  <div id="chatroom-header">
    <el-menu class="nav" :mode="check ? 'horizontal' : 'vertical'"
      text-color="#252c41" active-text-color="#409eff" default-active="/">
      <el-menu-item class="logo"><img src="../assets/logo.png"></el-menu-item>
      <el-menu-item class="title" index="/"><img src="../assets/logo.png">{{ title }}</el-menu-item>
    </el-menu>
  </div>
</template>

<script>
export default {
  name: 'ChatroomHeader',
  data() {
    return {
      title: process.env.VUE_APP_NAME,
      width: document.body.scrollWidth
    };
  },
  computed: {
    check() {
      return this.width > 768;
    }
  },
  created() {
    window.addEventListener('resize', () => {
      this.width = document.body.scrollWidth;
    });
  }
};
</script>
