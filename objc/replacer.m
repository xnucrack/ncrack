#import <Foundation/Foundation.h>
#include <objc/runtime.h>

@interface Replacer : NSObject{}
- (void)toReplace;
@end

@implementation Replacer

+ (void)load {
    Class thisKlass = [self class];
    Class toReplaceKlass = NSClassFromString(@"Cracker");

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
