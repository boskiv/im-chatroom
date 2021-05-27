<template>
  <el-row id="chatroom-main" type="flex" align="center">
    <el-dialog :visible="!joined" title="加入 ~" top="240px"
      :close-on-click-modal="false" :close-on-press-escape="false" :show-close="false">
      邮箱: <el-input class="email" v-model="user.email" placeholder="用于 Gravatar 头像" clearable/>
      <br><br>
      昵称: <el-input class="nick" v-model="user.nick" placeholder="看心情取一个就好" clearable/>
      <br><br>
      <el-button class="join" type="primary" plain @click="join">加入</el-button>
    </el-dialog>
    <JwChat-index v-model="input" :taleList="messages" :toolConfig="tool"
      height="600px" :config="{}" scrollType="scroll" @enter="send">
      <JwChat-rightbox :config="right"/>
    </JwChat-index>
  </el-row>
</template>

<script>
import moment from 'moment';
import md5 from 'md5';

export default {
  name: 'ChatroomMain',
  data() {
    return {
      ws: new WebSocket('ws://' + window.location.host + '/ws'),
      user: {
        email: '',
        nick: '',
      },
      joined: false,
      messages: [],
      input: '',
      tool: {
        showEmoji: false
      },
      right: {
        tip: '公告',
        listTip: '在线',
        list: []
      }
    }
  },
  created() {
    let self = this, flag = false;
    this.ws.onmessage = function (e) {
      if (!flag) {
        flag = true;
        let online = JSON.parse(e.data);
        for (let user of online.users) {
          self.right.list.push({
            name: user.nick,
            img: self.gravatar(user.email)
          });
        }
      } else {
        let message = JSON.parse(e.data);
        if (message.status == 1) {
          self.right.list.push({
            name: message.from.nick,
            img: self.gravatar(message.from.email)
          });
        } else if (message.status == -1) {
          let index = -1;
          for (let i = 0; i < self.right.list.length; i++) {
            if (self.right.list[i].img == self.gravatar(message.from.email)) {
              index = i;
              break;
            }
          }
          if (index >= 0) {
            let users = self.right.list;
            self.right.list = users.slice(0, index).concat(users.slice(index + 1, users.length));
          }
        } else {
          self.messages.push({
            date: self.timestamp(),
            text: {
              text: message.content
            },
            mine: message.from.email == self.user.email,
            name: message.from.nick,
            img: self.gravatar(message.from.email)
          });
        }
      }
    };
    this.ws.onmessage.onerror = function() {
      self.leave();
    };
    window.onbeforeunload = function() {
      self.leave();
    };
  },
  methods: {
    timestamp() {
      return moment(new Date()).format('YYYY/MM/DD HH:mm:ss');
    },
    gravatar(email) {
      return process.env.VUE_APP_GRAVATAR + '/' + md5(email);
    },
    join() {
      if (!this.user.email) {
        this.user.email = Math.random().toString(36).slice(-8);
      }
      if (!this.user.nick) {
        this.user.nick = '匿名';
      }
      this.ws.send(JSON.stringify({
        status: 1,
        from: this.user,
        content: ''
      }));
      this.joined = true;
    },
    send() {
      if (!this.input) {
        this.$message.info('消息不能为空呀 ~');
        return;
      }
      this.ws.send(JSON.stringify({
        status: 0,
        from: this.user,
        content: this.input
      }));
      this.input = '';
    },
    leave() {
      if (this.joined) {
        this.ws.send(JSON.stringify({
          status: -1,
          from: this.user,
          content: ''
        }));
      }
      this.ws.close();
      this.joined = false;
    }
  },
};
</script>
