#include <linux/module.h>
#define INCLUDE_VERMAGIC
#include <linux/build-salt.h>
#include <linux/elfnote-lto.h>
#include <linux/export-internal.h>
#include <linux/vermagic.h>
#include <linux/compiler.h>

#ifdef CONFIG_UNWINDER_ORC
#include <asm/orc_header.h>
ORC_HEADER;
#endif

BUILD_SALT;
BUILD_LTO_INFO;

MODULE_INFO(vermagic, VERMAGIC_STRING);
MODULE_INFO(name, KBUILD_MODNAME);

__visible struct module __this_module
__section(".gnu.linkonce.this_module") = {
	.name = KBUILD_MODNAME,
	.init = init_module,
#ifdef CONFIG_MODULE_UNLOAD
	.exit = cleanup_module,
#endif
	.arch = MODULE_ARCH_INIT,
};

#ifdef CONFIG_RETPOLINE
MODULE_INFO(retpoline, "Y");
#endif



static const struct modversion_info ____versions[]
__used __section("__versions") = {
	{ 0xff0f94f9, "single_open" },
	{ 0x67543840, "filp_open" },
	{ 0x92f3de2f, "kernel_read" },
	{ 0x3ef70737, "filp_close" },
	{ 0xbcab6ee6, "sscanf" },
	{ 0xf0fdf6cb, "__stack_chk_fail" },
	{ 0x2b73029d, "remove_proc_entry" },
	{ 0x944375db, "_totalram_pages" },
	{ 0xad16124d, "seq_printf" },
	{ 0x18ea85f5, "init_task" },
	{ 0x2f36395b, "init_user_ns" },
	{ 0xfde11b98, "from_kuid" },
	{ 0x87a21cb3, "__ubsan_handle_out_of_bounds" },
	{ 0xe85f2892, "seq_read" },
	{ 0xbdfb6dbb, "__fentry__" },
	{ 0x80efbd79, "proc_create" },
	{ 0x122c3a7e, "_printk" },
	{ 0x5b8239ca, "__x86_return_thunk" },
	{ 0x2fa5cadd, "module_layout" },
};

MODULE_INFO(depends, "");


MODULE_INFO(srcversion, "041B4001F1711ACFA2409A0");
