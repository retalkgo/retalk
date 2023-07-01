import { Button } from "./Button";
import { Input } from "./Input";
import { Textarea } from "./Textarea";

export default () => (
  <div>
    <div class="w-125 flex gap-3">
      <Input placeholder="昵称" />
      <Input placeholder="邮箱" />
      <Input placeholder="网站" />
    </div>
    <div class="mt-7.5">
      <Textarea/>
    </div>
    <div class="mt-7.5">
      <Button>
        <span>按钮</span>
      </Button>
    </div>
  </div>
);
