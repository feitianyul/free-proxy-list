**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-04-25 23:29:37 UTC（2026-04-26 07:29:37 UTC+8）

**代理总数：40**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 40 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 40 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 120.92.212.16:7890 | ✓ 917ms | ✓ 1166ms | ✓ 999ms | ✓ 1210ms | ✓ 942ms | http |
| 47.85.51.197:1080 | ✓ 1243ms | 否 | ✓ 1325ms | ✓ 1114ms | 否 | http |
| 20.127.128.70:8080 | ✓ 1239ms | 否 | ✓ 1389ms | 否 | ✓ 1913ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1120ms | ✓ 1409ms | ✓ 1835ms | http |
| 206.206.126.177:2412 | 否 | ✓ 1613ms | ✓ 1648ms | ✓ 1059ms | ✓ 844ms | http |
| 31.131.248.48:3129 | ✓ 1092ms | ✓ 1824ms | ✓ 1724ms | 否 | 否 | http |
| 120.92.108.86:7890 | ✓ 1216ms | 否 | ✓ 1173ms | ✓ 1725ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1788ms | 否 | ✓ 807ms | ✓ 1614ms | ✓ 1308ms | http |
| 113.160.132.26:8080 | ✓ 848ms | 否 | ✓ 1339ms | ✓ 1226ms | 否 | http |
| 47.84.59.16:1080 | ✓ 786ms | ✓ 1680ms | ✓ 1064ms | ✓ 1130ms | ✓ 908ms | http |
| 2.27.54.161:1080 | ✓ 982ms | 否 | ✓ 1197ms | ✓ 1956ms | ✓ 1701ms | http |
| 177.93.132.244:3128 | ✓ 1284ms | 否 | ✓ 899ms | 否 | ✓ 1771ms | http |
| 1.231.81.166:3128 | 否 | ✓ 988ms | ✓ 1065ms | ✓ 991ms | ✓ 863ms | http |
| 45.186.6.104:3128 | ✓ 1940ms | ✓ 1734ms | ✓ 1777ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 957ms | ✓ 1370ms | 否 | ✓ 1320ms | 否 | http |
| 114.237.77.245:1080 | ✓ 971ms | ✓ 1176ms | ✓ 987ms | 否 | ✓ 985ms | http |
| 36.141.21.200:7890 | ✓ 946ms | ✓ 1215ms | ✓ 973ms | ✓ 1275ms | ✓ 1001ms | http |
| 47.84.76.30:1080 | ✓ 1497ms | ✓ 1696ms | ✓ 884ms | ✓ 1121ms | ✓ 1622ms | http |
| 86.109.3.28:10002 | ✓ 485ms | ✓ 1140ms | ✓ 1276ms | 否 | 否 | http |
| 86.109.3.28:10010 | ✓ 424ms | ✓ 1110ms | 否 | ✓ 1030ms | 否 | http |
| 152.42.177.32:8888 | ✓ 1007ms | 否 | ✓ 1093ms | ✓ 1224ms | ✓ 1384ms | http |
| 130.61.174.200:1080 | ✓ 1869ms | ✓ 1850ms | 否 | ✓ 1833ms | 否 | http |
| 218.108.131.186:17890 | ✓ 888ms | ✓ 1035ms | ✓ 899ms | ✓ 1149ms | ✓ 940ms | http |
| 183.232.248.73:7890 | ✓ 944ms | ✓ 1158ms | ✓ 998ms | ✓ 1114ms | ✓ 906ms | http |
| 47.112.25.109:7890 | ✓ 998ms | 否 | ✓ 909ms | ✓ 1327ms | 否 | http |
| 80.92.204.47:1081 | ✓ 1749ms | ✓ 1501ms | ✓ 1527ms | 否 | ✓ 1459ms | http |
| 103.157.200.126:3128 | 否 | 否 | ✓ 1472ms | ✓ 1991ms | ✓ 1592ms | http |
| 221.156.27.160:3120 | 否 | 否 | ✓ 1226ms | ✓ 1185ms | ✓ 1481ms | http |
| 20.78.213.56:80 | ✓ 1507ms | ✓ 1265ms | ✓ 1452ms | ✓ 1744ms | 否 | http |
| 47.84.73.61:1080 | ✓ 1474ms | ✓ 1877ms | ✓ 994ms | ✓ 1127ms | ✓ 912ms | http |
| 34.101.184.164:3128 | ✓ 1614ms | 否 | 否 | ✓ 1586ms | ✓ 1196ms | http |
| 84.47.150.125:1080 | ✓ 1809ms | ✓ 1609ms | ✓ 1733ms | ✓ 1918ms | ✓ 1899ms | http |
| 117.236.124.166:3128 | ✓ 1294ms | 否 | ✓ 1184ms | ✓ 1925ms | 否 | http |
| 8.211.166.184:8081 | ✓ 1334ms | 否 | ✓ 763ms | ✓ 892ms | ✓ 717ms | http |
| 103.215.60.46:8097 | ✓ 1851ms | 否 | ✓ 1353ms | ✓ 1554ms | ✓ 1604ms | http |
| 103.126.238.13:8081 | ✓ 1852ms | 否 | ✓ 1578ms | ✓ 1917ms | ✓ 1523ms | http |
| 103.227.187.11:6090 | ✓ 1859ms | 否 | ✓ 1985ms | ✓ 1826ms | ✓ 1578ms | http |
| 61.52.131.172:8443 | ✓ 935ms | ✓ 1195ms | ✓ 992ms | ✓ 1159ms | ✓ 966ms | http |
| 163.44.126.97:3128 | 否 | 否 | ✓ 1978ms | ✓ 1516ms | ✓ 1060ms | http |
| 103.163.80.170:1080 | 否 | 否 | ✓ 1306ms | ✓ 1578ms | ✓ 1516ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
