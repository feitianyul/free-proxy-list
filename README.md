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

最后更新：2026-04-23 21:45:15 UTC（2026-04-24 05:45:15 UTC+8）

**代理总数：57**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 57 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 57 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 809ms | ✓ 1015ms | ✓ 758ms | ✓ 892ms | ✓ 710ms | http |
| 218.108.131.186:17890 | ✓ 798ms | ✓ 1039ms | ✓ 820ms | ✓ 1292ms | ✓ 910ms | http |
| 113.160.132.26:8080 | ✓ 1402ms | ✓ 1392ms | ✓ 1035ms | ✓ 1306ms | ✓ 947ms | http |
| 46.101.95.183:8888 | ✓ 880ms | 否 | 否 | ✓ 1769ms | ✓ 1473ms | http |
| 115.231.181.40:8128 | ✓ 1604ms | ✓ 1167ms | 否 | 否 | ✓ 901ms | http |
| 223.84.151.86:30005 | 否 | 否 | ✓ 1782ms | ✓ 1784ms | ✓ 1639ms | http |
| 8.211.166.184:8081 | ✓ 1212ms | ✓ 1146ms | ✓ 649ms | ✓ 896ms | ✓ 727ms | http |
| 35.225.22.61:80 | ✓ 423ms | ✓ 1494ms | ✓ 1816ms | ✓ 1426ms | ✓ 1076ms | http |
| 62.113.119.14:8080 | ✓ 827ms | ✓ 1589ms | ✓ 1156ms | ✓ 1638ms | ✓ 1205ms | http |
| 177.93.132.244:3128 | ✓ 1045ms | 否 | ✓ 823ms | 否 | ✓ 1830ms | http |
| 212.58.132.5:8888 | ✓ 1127ms | ✓ 1943ms | ✓ 1855ms | ✓ 1562ms | ✓ 1228ms | http |
| 120.92.108.86:7890 | ✓ 1156ms | 否 | ✓ 1904ms | ✓ 1708ms | ✓ 1353ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1084ms | ✓ 1385ms | ✓ 1137ms | http |
| 168.110.52.228:3128 | ✓ 1030ms | ✓ 1210ms | ✓ 954ms | ✓ 763ms | ✓ 717ms | http |
| 120.92.212.16:7890 | ✓ 1171ms | 否 | 否 | ✓ 1617ms | ✓ 1259ms | http |
| 45.140.147.155:1081 | ✓ 1201ms | 否 | ✓ 1735ms | 否 | ✓ 1829ms | http |
| 152.32.132.190:7890 | 否 | ✓ 1643ms | ✓ 1029ms | ✓ 1121ms | ✓ 1506ms | http |
| 47.85.51.197:1080 | 否 | ✓ 1521ms | ✓ 1496ms | 否 | ✓ 925ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1229ms | ✓ 1707ms | ✓ 1323ms | 否 | http |
| 162.240.154.26:3128 | 否 | ✓ 1494ms | ✓ 1526ms | ✓ 1601ms | 否 | http |
| 172.208.25.199:3128 | ✓ 1057ms | ✓ 1617ms | ✓ 1011ms | ✓ 1619ms | ✓ 1362ms | http |
| 20.127.128.70:8080 | ✓ 1098ms | 否 | ✓ 533ms | ✓ 1842ms | 否 | http |
| 91.99.15.45:2095 | ✓ 626ms | ✓ 1621ms | ✓ 1243ms | 否 | ✓ 1736ms | http |
| 201.144.20.238:3128 | ✓ 643ms | 否 | ✓ 924ms | ✓ 1268ms | 否 | http |
| 34.71.229.255:3128 | ✓ 642ms | ✓ 1859ms | ✓ 957ms | ✓ 1485ms | ✓ 1220ms | http |
| 34.96.238.40:8080 | ✓ 1184ms | ✓ 1048ms | ✓ 1130ms | ✓ 1242ms | ✓ 970ms | http |
| 45.76.207.177:40000 | ✓ 805ms | 否 | ✓ 1432ms | ✓ 1015ms | ✓ 738ms | http |
| 85.190.99.143:443 | ✓ 1275ms | 否 | ✓ 688ms | 否 | ✓ 1613ms | http |
| 150.107.140.238:3128 | ✓ 1841ms | 否 | ✓ 935ms | ✓ 1873ms | ✓ 1209ms | http |
| 121.230.8.249:1080 | ✓ 1250ms | ✓ 1981ms | ✓ 1154ms | 否 | 否 | http |
| 208.87.243.199:7878 | ✓ 934ms | ✓ 1108ms | ✓ 1239ms | 否 | 否 | http |
| 60.249.94.208:3128 | ✓ 974ms | ✓ 928ms | ✓ 697ms | ✓ 903ms | ✓ 1496ms | http |
| 210.45.76.58:42992 | ✓ 1045ms | ✓ 1455ms | ✓ 1523ms | ✓ 1414ms | ✓ 1161ms | http |
| 152.70.91.193:40000 | ✓ 1628ms | 否 | 否 | ✓ 1991ms | ✓ 1867ms | http |
| 114.237.77.239:1080 | ✓ 923ms | ✓ 1230ms | ✓ 1012ms | ✓ 1177ms | ✓ 905ms | http |
| 121.230.8.97:1080 | 否 | ✓ 1323ms | ✓ 967ms | ✓ 1351ms | 否 | http |
| 121.230.8.62:1080 | 否 | ✓ 1275ms | ✓ 1246ms | ✓ 1626ms | ✓ 1339ms | http |
| 124.121.2.131:8080 | ✓ 1829ms | 否 | ✓ 1706ms | ✓ 1529ms | ✓ 1463ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1950ms | ✓ 1296ms | ✓ 1450ms | ✓ 1184ms | http |
| 2.83.243.148:7777 | ✓ 923ms | ✓ 1953ms | ✓ 1672ms | 否 | 否 | http |
| 89.208.106.138:10808 | ✓ 1127ms | ✓ 1724ms | 否 | 否 | ✓ 1613ms | http |
| 178.159.94.76:3128 | ✓ 821ms | ✓ 1717ms | ✓ 1254ms | 否 | ✓ 1725ms | http |
| 42.101.8.101:8888 | ✓ 1196ms | 否 | ✓ 1158ms | ✓ 1343ms | ✓ 1155ms | http |
| 130.61.174.200:1080 | ✓ 1086ms | 否 | ✓ 1317ms | ✓ 1878ms | ✓ 1674ms | http |
| 210.48.154.94:80 | 否 | 否 | ✓ 1236ms | ✓ 1135ms | ✓ 1032ms | http |
| 128.199.116.219:9090 | ✓ 814ms | 否 | ✓ 1878ms | ✓ 1589ms | ✓ 1238ms | http |
| 34.85.118.216:3128 | ✓ 1663ms | ✓ 1322ms | ✓ 1435ms | ✓ 969ms | ✓ 702ms | http |
| 84.47.150.125:1080 | ✓ 886ms | 否 | ✓ 1415ms | 否 | ✓ 1861ms | http |
| 121.230.9.160:1080 | ✓ 1030ms | ✓ 1341ms | ✓ 1136ms | ✓ 1454ms | ✓ 1167ms | http |
| 20.120.225.109:3128 | ✓ 999ms | ✓ 1320ms | ✓ 954ms | ✓ 1308ms | ✓ 933ms | http |
| 137.59.47.73:3128 | ✓ 1918ms | 否 | 否 | ✓ 1945ms | ✓ 1640ms | http |
| 114.237.77.202:1080 | ✓ 1024ms | ✓ 1178ms | ✓ 893ms | ✓ 1247ms | ✓ 914ms | http |
| 218.77.106.10:10150 | ✓ 995ms | ✓ 1271ms | ✓ 970ms | ✓ 1280ms | ✓ 944ms | http |
| 121.230.9.209:1080 | 否 | ✓ 1543ms | ✓ 1020ms | ✓ 1526ms | ✓ 1410ms | http |
| 34.101.184.164:3128 | ✓ 1717ms | 否 | ✓ 1544ms | ✓ 1808ms | ✓ 1004ms | http |
| 45.186.6.104:3128 | ✓ 1962ms | ✓ 1915ms | ✓ 1763ms | 否 | 否 | http |
| 45.153.231.229:8080 | ✓ 1016ms | ✓ 1773ms | ✓ 1734ms | 否 | 否 | http |

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
