import Vue from 'vue'
import 'element-ui/lib/theme-chalk/base.css';
import { Button, Form, FormItem, Input, Message, Container,
  Header, Aside, Main, Footer, Menu, Submenu, MenuItem, PageHeader,
  Card, Row, Col, Table, TableColumn, Switch, Tooltip, Pagination,
  Dialog, MessageBox, Dropdown, DropdownItem, DropdownMenu, Alert,
  Step, Steps, Radio, RadioGroup, Select, Option, Popconfirm, Popover } from 'element-ui'

Vue.use(Button);Vue.use(Form);Vue.use(FormItem);Vue.use(Input);Vue.use(Container);Vue.use(Header);
Vue.use(Aside);Vue.use(Main);Vue.use(Footer);Vue.use(Menu);Vue.use(Submenu);Vue.use(MenuItem);
Vue.use(PageHeader);Vue.use(Card);Vue.use(Row);Vue.use(Col);Vue.use(Table);Vue.use(TableColumn);
Vue.use(Switch);Vue.use(Tooltip);Vue.use(Pagination);Vue.use(Dialog);Vue.use(Dropdown);Vue.use(DropdownItem);
Vue.use(DropdownMenu);Vue.use(Alert);Vue.use(Step);Vue.use(Steps);Vue.use(Radio);Vue.use(RadioGroup);
Vue.use(Select);Vue.use(Option);Vue.use(Popconfirm);Vue.use(Popover);
Vue.prototype.$message = Message;
Vue.prototype.$confirm = MessageBox.confirm;
