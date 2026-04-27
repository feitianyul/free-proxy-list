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

最后更新：2026-04-27 07:23:39 UTC（2026-04-27 15:23:39 UTC+8）

**代理总数：49**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 49 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 49 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | ✓ 814ms | ✓ 1073ms | ✓ 1881ms | ✓ 1149ms | ✓ 920ms | http |
| 38.180.192.119:3128 | ✓ 680ms | ✓ 722ms | 否 | ✓ 761ms | 否 | http |
| 211.95.152.50:45046 | ✓ 1133ms | ✓ 1330ms | ✓ 1106ms | ✓ 1383ms | ✓ 1077ms | http |
| 59.46.216.131:30001 | ✓ 1025ms | 否 | ✓ 1936ms | ✓ 1342ms | ✓ 1062ms | http |
| 217.76.245.80:999 | ✓ 1076ms | 否 | ✓ 1140ms | ✓ 1638ms | ✓ 1293ms | http |
| 34.101.184.164:3128 | ✓ 796ms | 否 | ✓ 876ms | ✓ 1228ms | ✓ 935ms | http |
| 113.160.132.26:8080 | ✓ 1664ms | 否 | ✓ 919ms | ✓ 1251ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1617ms | 否 | ✓ 1495ms | 否 | ✓ 1691ms | http |
| 120.92.108.86:7890 | ✓ 1150ms | 否 | 否 | ✓ 1762ms | ✓ 1376ms | http |
| 47.85.51.197:1080 | 否 | ✓ 1273ms | ✓ 453ms | ✓ 1166ms | ✓ 931ms | http |
| 38.244.54.190:31168 | ✓ 1869ms | 否 | ✓ 869ms | ✓ 1216ms | ✓ 518ms | http |
| 8.219.195.129:1080 | ✓ 1344ms | 否 | ✓ 747ms | ✓ 1038ms | ✓ 836ms | http |
| 160.250.4.245:1 | 否 | 否 | ✓ 1615ms | ✓ 1348ms | ✓ 1112ms | http |
| 86.104.74.110:1081 | ✓ 1285ms | ✓ 1893ms | ✓ 795ms | 否 | ✓ 1456ms | http |
| 62.113.119.14:8080 | ✓ 1297ms | 否 | ✓ 1716ms | ✓ 1915ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1692ms | ✓ 1277ms | ✓ 1136ms | ✓ 955ms | ✓ 1004ms | http |
| 46.101.95.183:8888 | ✓ 1157ms | 否 | ✓ 1170ms | 否 | ✓ 1467ms | http |
| 47.101.159.19:8899 | ✓ 867ms | ✓ 1098ms | ✓ 891ms | ✓ 1137ms | ✓ 902ms | http |
| 167.71.196.178:80 | ✓ 933ms | 否 | ✓ 1000ms | ✓ 1054ms | ✓ 828ms | http |
| 8.137.112.117:3128 | ✓ 991ms | ✓ 1399ms | ✓ 1066ms | ✓ 1311ms | ✓ 1005ms | http |
| 103.157.200.126:3128 | ✓ 1730ms | 否 | 否 | ✓ 1939ms | ✓ 1464ms | http |
| 103.50.25.59:8080 | ✓ 1766ms | 否 | ✓ 1209ms | ✓ 1525ms | ✓ 1404ms | http |
| 121.43.196.213:8222 | ✓ 981ms | ✓ 1023ms | ✓ 885ms | ✓ 1086ms | ✓ 889ms | http |
| 159.223.225.118:8888 | ✓ 1716ms | 否 | ✓ 732ms | 否 | ✓ 1505ms | http |
| 121.230.8.91:1080 | 否 | 否 | ✓ 1200ms | ✓ 1424ms | ✓ 1230ms | http |
| 47.84.73.61:1080 | ✓ 745ms | 否 | ✓ 849ms | ✓ 1054ms | ✓ 876ms | http |
| 101.32.244.83:8080 | ✓ 977ms | ✓ 1535ms | ✓ 923ms | ✓ 1353ms | ✓ 1265ms | http |
| 121.43.196.210:8222 | ✓ 950ms | ✓ 1031ms | ✓ 920ms | ✓ 1091ms | ✓ 890ms | http |
| 114.55.226.123:10086 | ✓ 1178ms | 否 | ✓ 1163ms | ✓ 1373ms | 否 | http |
| 121.230.8.111:1080 | 否 | 否 | ✓ 1765ms | ✓ 1794ms | ✓ 1210ms | http |
| 38.244.52.152:26678 | ✓ 361ms | ✓ 716ms | ✓ 716ms | ✓ 861ms | ✓ 728ms | http |
| 121.230.8.136:1080 | ✓ 1119ms | ✓ 1190ms | ✓ 989ms | ✓ 1740ms | ✓ 1827ms | http |
| 120.92.212.16:8890 | ✓ 1036ms | ✓ 1160ms | ✓ 1201ms | ✓ 1376ms | 否 | http |
| 120.92.212.16:7890 | 否 | ✓ 1481ms | ✓ 1961ms | ✓ 1153ms | 否 | http |
| 77.110.119.136:3128 | ✓ 1416ms | 否 | ✓ 629ms | ✓ 1828ms | 否 | http |
| 47.84.59.16:1080 | ✓ 905ms | ✓ 1714ms | ✓ 720ms | ✓ 1049ms | ✓ 862ms | http |
| 47.84.76.30:1080 | ✓ 927ms | 否 | ✓ 737ms | ✓ 1080ms | ✓ 820ms | http |
| 45.140.147.82:1081 | ✓ 1099ms | ✓ 1803ms | ✓ 1242ms | ✓ 1450ms | ✓ 1821ms | http |
| 45.140.147.82:1082 | ✓ 1093ms | ✓ 1594ms | ✓ 1452ms | ✓ 1454ms | ✓ 1882ms | http |
| 177.93.132.244:3128 | ✓ 830ms | 否 | ✓ 790ms | 否 | ✓ 1953ms | http |
| 34.96.238.40:8080 | ✓ 1688ms | 否 | ✓ 1228ms | ✓ 1495ms | 否 | http |
| 150.249.255.91:3128 | 否 | 否 | ✓ 557ms | ✓ 1336ms | ✓ 661ms | http |
| 45.59.122.132:80 | ✓ 912ms | 否 | ✓ 1399ms | ✓ 1837ms | ✓ 1514ms | http |
| 101.132.61.121:8888 | ✓ 1273ms | 否 | ✓ 1483ms | 否 | ✓ 1351ms | http |
| 61.52.131.172:8443 | ✓ 923ms | ✓ 1294ms | ✓ 962ms | ✓ 1300ms | ✓ 1008ms | http |
| 106.44.155.90:2222 | ✓ 1518ms | ✓ 1898ms | 否 | 否 | ✓ 1745ms | http |
| 223.16.170.103:80 | 否 | 否 | ✓ 1997ms | ✓ 1073ms | ✓ 1085ms | http |
| 103.39.51.207:8080 | ✓ 1361ms | 否 | ✓ 1678ms | ✓ 1599ms | ✓ 1945ms | http |
| 118.31.1.154:80 | 否 | ✓ 1114ms | ✓ 889ms | ✓ 1119ms | 否 | http |

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
