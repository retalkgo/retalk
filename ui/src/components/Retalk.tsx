import { Button } from "./Button";
import { Input } from "./Input";
import { Textarea } from "./Textarea";

export default () => (
  <div>
    <div class="w-125 flex gap-3">
      <Input placeholder="A" />
      <Input placeholder="B" />
      <Input placeholder="C" />
    </div>
    <Input />
    <Textarea />
    <Button>
      <span>按钮</span>
    </Button>
  </div>
);
