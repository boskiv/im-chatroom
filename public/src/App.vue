<template>
  <div id="app">
    <chatroom-header/>
    <chatroom-main/>
    <chatroom-footer/>
  </div>
</template>

<script>
import ChatroomHeader from './components/ChatroomHeader';
import ChatroomMain from './components/ChatroomMain';
import ChatroomFooter from './components/ChatroomFooter';

export default {
  name: 'App',
  components: {
    ChatroomHeader,
    ChatroomMain,
    ChatroomFooter
  }
}
</script>
