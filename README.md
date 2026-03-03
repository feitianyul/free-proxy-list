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

最后更新：2026-03-03 09:42:45 UTC（2026-03-03 17:42:45 UTC+8）

**代理总数：44**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 44 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 44 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 162.240.154.26:3128 | ✓ 822ms | 否 | 否 | ✓ 1619ms | ✓ 1321ms | http |
| 142.171.85.32:1080 | 否 | ✓ 1395ms | ✓ 1217ms | ✓ 1204ms | ✓ 1428ms | http |
| 103.84.95.54:7890 | ✓ 837ms | 否 | ✓ 1903ms | ✓ 985ms | 否 | http |
| 3.213.157.4:3128 | 否 | 否 | ✓ 1414ms | ✓ 1776ms | ✓ 1035ms | http |
| 115.231.181.40:8128 | ✓ 902ms | ✓ 1210ms | ✓ 1069ms | ✓ 1140ms | ✓ 888ms | http |
| 35.234.17.221:8080 | 否 | ✓ 1168ms | ✓ 1333ms | ✓ 904ms | 否 | http |
| 205.209.118.30:3138 | 否 | 否 | ✓ 1215ms | ✓ 1335ms | ✓ 1037ms | http |
| 47.77.180.205:1080 | ✓ 894ms | 否 | ✓ 110ms | ✓ 701ms | 否 | http |
| 47.101.149.27:9010 | 否 | ✓ 1514ms | ✓ 1343ms | 否 | ✓ 1065ms | http |
| 81.70.169.194:80 | 否 | 否 | ✓ 1005ms | ✓ 1258ms | ✓ 1061ms | http |
| 121.230.9.161:1080 | ✓ 1020ms | ✓ 1455ms | ✓ 1232ms | ✓ 1474ms | 否 | http |
| 101.43.255.96:80 | ✓ 1816ms | 否 | ✓ 1300ms | ✓ 1201ms | ✓ 986ms | http |
| 120.92.212.16:8890 | ✓ 943ms | 否 | ✓ 1788ms | ✓ 1242ms | ✓ 1778ms | http |
| 166.0.192.117:8888 | ✓ 370ms | 否 | ✓ 1209ms | ✓ 1958ms | ✓ 1400ms | http |
| 106.14.203.63:3333 | ✓ 1503ms | 否 | ✓ 885ms | ✓ 1119ms | ✓ 1652ms | http |
| 61.72.110.54:3128 | ✓ 980ms | ✓ 1673ms | 否 | ✓ 1813ms | 否 | http |
| 138.124.53.25:7443 | ✓ 1616ms | ✓ 1976ms | 否 | ✓ 1973ms | 否 | http |
| 160.250.4.13:1 | 否 | 否 | ✓ 1768ms | ✓ 1613ms | ✓ 1589ms | http |
| 8.219.97.248:80 | ✓ 1782ms | 否 | ✓ 1561ms | ✓ 1553ms | 否 | http |
| 14.56.177.44:3128 | ✓ 1302ms | ✓ 1433ms | ✓ 1114ms | ✓ 996ms | ✓ 771ms | http |
| 120.92.212.16:7890 | ✓ 1216ms | ✓ 1841ms | ✓ 1535ms | ✓ 1499ms | ✓ 957ms | http |
| 150.107.140.238:3128 | ✓ 1855ms | 否 | ✓ 952ms | ✓ 1181ms | ✓ 1030ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1052ms | ✓ 1157ms | ✓ 1239ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1040ms | ✓ 1580ms | ✓ 1796ms | http |
| 91.233.223.147:3128 | ✓ 1331ms | 否 | ✓ 1762ms | 否 | ✓ 1659ms | http |
| 5.75.196.26:40000 | ✓ 590ms | 否 | ✓ 1680ms | ✓ 1933ms | 否 | http |
| 38.180.226.51:3128 | ✓ 1487ms | 否 | ✓ 1280ms | ✓ 1938ms | 否 | http |
| 91.193.240.157:9877 | ✓ 993ms | 否 | ✓ 1242ms | 否 | ✓ 1843ms | http |
| 106.14.205.114:483 | ✓ 989ms | ✓ 1055ms | ✓ 1173ms | ✓ 1081ms | ✓ 822ms | http |
| 181.78.194.249:999 | ✓ 1409ms | ✓ 1960ms | 否 | 否 | ✓ 1875ms | http |
| 45.136.198.40:3128 | ✓ 827ms | 否 | ✓ 1993ms | 否 | ✓ 1920ms | http |
| 45.140.147.82:1081 | ✓ 928ms | 否 | ✓ 1578ms | 否 | ✓ 1450ms | http |
| 157.0.142.246:10058 | ✓ 1011ms | ✓ 1234ms | ✓ 947ms | ✓ 1276ms | 否 | http |
| 144.208.127.181:3128 | ✓ 907ms | 否 | ✓ 1227ms | ✓ 1737ms | ✓ 1513ms | http |
| 103.215.36.88:18574 | 否 | 否 | ✓ 1499ms | ✓ 1550ms | ✓ 1964ms | http |
| 202.129.206.239:3128 | ✓ 1557ms | 否 | ✓ 1579ms | ✓ 1522ms | ✓ 1779ms | http |
| 103.39.51.190:8080 | ✓ 1933ms | 否 | 否 | ✓ 1784ms | ✓ 1556ms | http |
| 47.95.231.180:8084 | ✓ 897ms | ✓ 1173ms | ✓ 950ms | ✓ 1158ms | ✓ 910ms | http |
| 91.107.148.58:53967 | ✓ 637ms | 否 | 否 | ✓ 1752ms | ✓ 1063ms | http |
| 121.230.9.205:1080 | 否 | ✓ 1233ms | ✓ 918ms | 否 | ✓ 1094ms | http |
| 160.238.65.9:3128 | ✓ 1179ms | 否 | 否 | ✓ 1605ms | ✓ 1158ms | http |
| 103.215.36.88:16894 | 否 | 否 | ✓ 1249ms | ✓ 1602ms | ✓ 1117ms | http |
| 103.215.36.88:17100 | ✓ 1582ms | 否 | ✓ 1206ms | ✓ 1502ms | ✓ 1084ms | http |
| 223.16.170.103:80 | ✓ 1932ms | 否 | ✓ 842ms | ✓ 1054ms | 否 | http |

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
