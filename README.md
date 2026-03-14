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

最后更新：2026-03-14 19:39:47 UTC（2026-03-15 03:39:47 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.145.218.82:8443 | ✓ 949ms | ✓ 659ms | ✓ 1077ms | ✓ 840ms | ✓ 620ms | http |
| 101.47.73.135:3128 | ✓ 825ms | 否 | ✓ 1437ms | ✓ 1179ms | ✓ 1009ms | http |
| 113.160.132.26:8080 | ✓ 1460ms | ✓ 1395ms | ✓ 1724ms | ✓ 1284ms | ✓ 1331ms | http |
| 150.230.249.50:1080 | 否 | 否 | ✓ 1163ms | ✓ 935ms | ✓ 837ms | http |
| 45.167.124.52:8080 | ✓ 1870ms | ✓ 1945ms | 否 | 否 | ✓ 1668ms | http |
| 45.140.147.155:1081 | ✓ 690ms | ✓ 1657ms | ✓ 1566ms | 否 | ✓ 1407ms | http |
| 217.76.245.80:999 | ✓ 702ms | ✓ 1541ms | ✓ 1220ms | ✓ 1461ms | ✓ 1351ms | http |
| 59.46.216.131:30001 | ✓ 1073ms | ✓ 1416ms | 否 | ✓ 1494ms | ✓ 995ms | http |
| 81.70.169.194:80 | ✓ 1141ms | ✓ 1363ms | ✓ 980ms | ✓ 1179ms | ✓ 1055ms | http |
| 101.43.255.96:80 | ✓ 1142ms | 否 | ✓ 1124ms | ✓ 1213ms | ✓ 1151ms | http |
| 205.209.118.30:3138 | 否 | 否 | ✓ 963ms | ✓ 1550ms | ✓ 1009ms | http |
| 45.136.131.28:8447 | ✓ 1020ms | ✓ 1677ms | ✓ 1877ms | 否 | 否 | http |
| 91.233.223.147:3128 | ✓ 1066ms | 否 | ✓ 1646ms | 否 | ✓ 1596ms | http |
| 43.167.227.161:1080 | ✓ 1501ms | ✓ 779ms | ✓ 466ms | ✓ 826ms | ✓ 593ms | http |
| 121.230.8.144:1080 | ✓ 1290ms | 否 | ✓ 1380ms | ✓ 1454ms | 否 | http |
| 35.225.22.61:80 | ✓ 1101ms | 否 | ✓ 1095ms | ✓ 1372ms | 否 | http |
| 45.136.130.223:8443 | ✓ 321ms | ✓ 606ms | ✓ 464ms | ✓ 823ms | ✓ 513ms | http |
| 38.145.203.162:8443 | ✓ 165ms | ✓ 649ms | ✓ 743ms | ✓ 713ms | ✓ 479ms | http |
| 38.145.208.96:8443 | ✓ 147ms | ✓ 629ms | ✓ 783ms | ✓ 721ms | ✓ 476ms | http |
| 38.145.208.98:8443 | ✓ 169ms | ✓ 632ms | ✓ 757ms | ✓ 721ms | ✓ 487ms | http |
| 38.145.208.93:8443 | ✓ 166ms | ✓ 602ms | ✓ 791ms | ✓ 713ms | ✓ 499ms | http |
| 38.145.203.245:8443 | ✓ 154ms | ✓ 597ms | ✓ 800ms | ✓ 724ms | ✓ 480ms | http |
| 38.145.203.161:8443 | ✓ 166ms | ✓ 599ms | ✓ 792ms | ✓ 716ms | ✓ 489ms | http |
| 38.145.208.94:8443 | ✓ 131ms | ✓ 1169ms | ✓ 244ms | ✓ 752ms | ✓ 483ms | http |
| 38.145.203.246:8443 | ✓ 142ms | ✓ 613ms | ✓ 803ms | ✓ 761ms | ✓ 483ms | http |
| 38.145.208.97:8443 | ✓ 161ms | ✓ 613ms | ✓ 785ms | ✓ 750ms | ✓ 508ms | http |
| 38.145.208.138:8447 | ✓ 168ms | ✓ 723ms | ✓ 816ms | ✓ 844ms | ✓ 543ms | http |
| 85.198.96.242:3128 | ✓ 723ms | 否 | ✓ 1762ms | ✓ 1770ms | ✓ 1449ms | http |
| 5.102.109.41:999 | ✓ 1663ms | ✓ 1378ms | 否 | ✓ 1242ms | 否 | http |
| 172.212.68.37:3128 | ✓ 530ms | ✓ 1624ms | ✓ 739ms | ✓ 1791ms | ✓ 1664ms | http |
| 38.145.203.135:8443 | ✓ 616ms | ✓ 1270ms | ✓ 104ms | ✓ 770ms | ✓ 513ms | http |
| 165.227.5.10:8888 | ✓ 127ms | 否 | ✓ 585ms | ✓ 1337ms | ✓ 695ms | http |
| 45.149.92.147:5001 | ✓ 620ms | 否 | ✓ 619ms | ✓ 781ms | ✓ 626ms | http |
| 194.5.212.40:8080 | ✓ 1141ms | ✓ 1999ms | ✓ 1339ms | ✓ 1545ms | ✓ 1658ms | http |
| 34.101.184.164:3128 | ✓ 1056ms | 否 | ✓ 778ms | ✓ 1139ms | ✓ 1156ms | http |
| 101.32.244.83:8080 | ✓ 1072ms | ✓ 1789ms | ✓ 1042ms | ✓ 1356ms | ✓ 1241ms | http |
| 121.43.196.213:8222 | ✓ 935ms | ✓ 1042ms | ✓ 867ms | ✓ 1131ms | ✓ 908ms | http |
| 121.43.196.210:8222 | ✓ 949ms | ✓ 1025ms | ✓ 923ms | ✓ 1193ms | ✓ 893ms | http |
| 114.55.226.123:10086 | ✓ 1032ms | ✓ 1665ms | ✓ 1129ms | ✓ 1290ms | ✓ 1000ms | http |
| 62.60.177.204:34094 | 否 | ✓ 1708ms | ✓ 1224ms | 否 | ✓ 1031ms | http |
| 107.174.55.123:7878 | ✓ 1321ms | ✓ 1672ms | ✓ 901ms | ✓ 686ms | ✓ 689ms | http |
| 121.40.231.103:7890 | ✓ 1561ms | ✓ 1380ms | ✓ 977ms | ✓ 1214ms | ✓ 968ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1138ms | ✓ 1004ms | ✓ 841ms | http |
| 14.225.212.37:7890 | ✓ 791ms | ✓ 1701ms | 否 | ✓ 1678ms | 否 | http |
| 45.136.131.42:8447 | ✓ 673ms | ✓ 1399ms | ✓ 932ms | ✓ 1559ms | 否 | http |
| 150.249.255.91:3128 | ✓ 576ms | ✓ 1389ms | ✓ 485ms | ✓ 904ms | ✓ 659ms | http |
| 103.84.95.54:7890 | ✓ 1045ms | 否 | ✓ 648ms | ✓ 828ms | ✓ 994ms | http |
| 120.92.212.16:7890 | ✓ 955ms | ✓ 1213ms | ✓ 955ms | ✓ 1246ms | ✓ 970ms | http |
| 45.136.198.40:3128 | ✓ 1220ms | 否 | ✓ 1117ms | 否 | ✓ 1935ms | http |
| 86.53.183.16:1080 | ✓ 1084ms | ✓ 1581ms | ✓ 1521ms | 否 | 否 | http |
| 178.236.245.59:3128 | ✓ 1458ms | 否 | ✓ 1088ms | ✓ 1825ms | 否 | http |
| 116.80.96.107:3172 | ✓ 1475ms | 否 | ✓ 1505ms | 否 | ✓ 1811ms | http |
| 38.145.208.95:8443 | ✓ 614ms | ✓ 743ms | ✓ 729ms | ✓ 678ms | ✓ 578ms | http |
| 38.145.208.99:8443 | ✓ 614ms | ✓ 743ms | ✓ 731ms | ✓ 679ms | ✓ 593ms | http |
| 45.136.131.30:8447 | 否 | ✓ 1206ms | ✓ 91ms | ✓ 677ms | ✓ 492ms | http |
| 120.92.212.16:8890 | ✓ 1748ms | 否 | 否 | ✓ 1254ms | ✓ 1992ms | http |
| 116.80.65.77:3172 | 否 | 否 | ✓ 1470ms | ✓ 1796ms | ✓ 1629ms | http |
| 103.39.51.190:8080 | ✓ 1784ms | 否 | 否 | ✓ 1294ms | ✓ 1327ms | http |
| 101.43.127.100:8877 | ✓ 1270ms | ✓ 1147ms | ✓ 1794ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 865ms | ✓ 1181ms | ✓ 1052ms | ✓ 1178ms | ✓ 952ms | http |
| 88.80.150.82:8080 | ✓ 1363ms | ✓ 1975ms | 否 | 否 | ✓ 1844ms | https |
| 2.56.122.146:10808 | ✓ 1290ms | ✓ 1365ms | ✓ 1472ms | ✓ 1980ms | ✓ 1557ms | http |

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
