#import <Foundation/Foundation.h>
#include <objc/runtime.h>

#define CLASS_NAME "Replacer"
#define TARGET_CLASS "Cracker"

@interface Replacer : NSObject{}
- (void)toReplace;
@end

@implementation Replacer

+ (void)load {
    Class thisKlass = objc_getClass(CLASS_NAME);
    Class toReplaceKlass = objc_getClass(TARGET_CLASS);

    SEL selReplace = @selector(toReplace);

    Method replaceMethod = class_getInstanceMethod(thisKlass, selReplace);
    Method toReplaceMethod = class_getInstanceMethod(toReplaceKlass, selReplace);

    IMP replaceImplementation = method_getImplementation(replaceMethod);
    method_setImplementation(toReplaceMethod, replaceImplementation);
}

- (void)toReplace {
    NSLog(@"I am replaced");
}

@end
