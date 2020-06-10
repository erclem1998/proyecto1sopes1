#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/module.h>
#include <linux/list.h>
#include <linux/types.h>
#include <linux/slab.h>
#include <linux/sched.h>
#include <linux/string.h>
#include <linux/fs.h>
#include <linux/seq_file.h>
#include <linux/proc_fs.h>
#include <linux/mm.h>
#include <linux/sched/signal.h>

void recorrerHijos(struct seq_file *sf, struct task_struct *ts, int contador)
{
    struct task_struct *root,*task = current;
    while (root->pid != 1)
    {
        root = root->parent;
    }
    list_for_each(list, &ts->children)
    {
        task = list_entry(list, struct task_struct, sibling);
        pr_info("== %s [%d]\n", root->comm, root->pid);
        recorrerHijos(sf, task, contador + 1);
        recorrerHijos(sf, root, 1);
    }
}

/*void procs_info_print(void)
{
    struct task_struct *task_list;
    size_t process_counter = 0;
    for_each_process(task_list)
    {
        pr_info("== %s [%d]\n", task_list->comm, task_list->pid);
        ++process_counter;
    }
    printk(KERN_INFO "== Number of process: %zu\n", process_counter);
}*/

int init_module(void)
{
    printk(KERN_INFO "[ INIT ==\n");

    //procs_info_print();

    return 0;
}

void cleanup_module(void)
{
    printk(KERN_INFO "== CLEANUP ]\n");
}

MODULE_LICENSE("MIT");
