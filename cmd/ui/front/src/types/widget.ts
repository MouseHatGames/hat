export type Widget = WidgetOnOff | WidgetText | WidgetGroup | WidgetOptions;

export interface WidgetBase {
    title: string;
    description?: string;
}

export interface WidgetOnOff extends WidgetBase {
    type: "onoff";
    value: boolean;
}

export interface WidgetText extends WidgetBase {
    type: "text";
    value?: string;
    placeholder?: string;
    big?: boolean;
}

export interface WidgetGroup extends WidgetBase {
    type: "group";
    children: Widget[];
}

export interface WidgetOptions extends WidgetBase {
    type: "options";
    options: string[];
    value?: number;
}